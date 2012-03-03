package otk

import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"hash"
	"io"
)

var (
	EmptyKey       = errors.New("Keys cannot be empty")
	InvalidPayload = errors.New("Payload is invalid")
)

type CipherSuite int8

const (
	version byte = 0x01
	AES256  byte = 0x01 + iota
	AES128
)

// An Encoder converts an OpenToken into a binary blob suitable for
// transmission over the wire.
type Encoder struct {
	suite  byte
	key    []byte
	writer io.Writer
}

// Instantiate a new Encoder that uses the specified suite and key to
// generate an encrypted binary blob to write to writer.
func NewEncoder(suite byte, key []byte, writer io.Writer) *Encoder {
	return &Encoder{suite: suite, key: key, writer: writer}
}

// Encode the OpenToken into a base64-encoded binary blob using the cipher
// suite and key, and write it to the writer.
func (encoder *Encoder) Encode(token OpenToken) error {
	payload, err := encoder.generatePayload(token)
	mac := encoder.generateHMAC(payload)
	encryptedPayload, err := encoder.encryptPayload(payload)
	encoder.generateBinaryToken(mac, encryptedPayload)
	return err
}

// Convert an OpenToken to a payload according to the rules defined in section
// 5 of the specification.
func (encoder *Encoder) generatePayload(token OpenToken) ([]byte, error) {
	// FIXME We need to validate the keys and values we write out, to make
	// sure they comply with the accepted characters defined in the
	// specification. -jkakar
	buffer := bytes.Buffer{}
	for key, values := range token {
		if key == "" {
			return nil, EmptyKey
		}
		for _, value := range values {
			buffer.WriteString(key)
			buffer.WriteString(" = ")
			buffer.WriteString(value)
			buffer.WriteString("\r\n")
		}
	}
	return buffer.Bytes(), nil
}

// Generate an HMAC using the specified cipher suite and payload according to
// the rules defined in section 3.1 of the specification.
func (encoder *Encoder) generateHMAC(payload []byte) hash.Hash {
	buffer := new(bytes.Buffer)
	buffer.WriteByte(version)
	buffer.WriteByte(encoder.suite)
	binary.Write(buffer, binary.BigEndian, len(payload))
	mac := hmac.New(sha1.New, buffer.Bytes())
	mac.Write(payload)
	return mac
}

// Compress the payload using DEFLATE compression and encrypt it using the
// specified cipher suite.
func (encoder *Encoder) encryptPayload(payload []byte) ([]byte, error) {
	// Compress the payload.
	compressedPayload := new(bytes.Buffer)
	compressor, err := flate.NewWriter(compressedPayload, -1)
	if err != nil {
		return nil, err
	}
	_, err = compressor.Write(payload)
	if err != nil {
		return nil, err
	}

	// Prepare the cipher and initialization vector.
	block, err := aes.NewCipher(encoder.key)
	if err != nil {
		return nil, err
	}
	// FIXME Generate a real IV here.
	blockMode := cipher.NewCBCEncrypter(block, encoder.key)

	// Apply padding to the compressed payload to ensure that it aligns with
	// the block size required by the cipher.
	n := compressedPayload.Len()
	for i := 0; i < blockMode.BlockSize() - (n % blockMode.BlockSize()); i++ {
		compressedPayload.WriteByte(0x0)
	}

	// Encrypt the compressed payload using the selected cipher.
	encryptedPayload := compressedPayload.Bytes()
	blockMode.CryptBlocks(encryptedPayload, encryptedPayload)
	return encryptedPayload, nil
}

// Generate the binary token from the HMAC and the encrypted payload.
func (encoder *Encoder) generateBinaryToken(mac hash.Hash, encryptedPayload []byte) []byte {
	token := new(bytes.Buffer)
	token.WriteString("OTK")
	token.WriteByte(version)
	token.WriteByte(encoder.suite)
	token.Write(mac.Sum(nil))
	token.WriteByte(0x0) // IV length
	token.WriteByte(0x0) // Key info length
	binary.Write(token, binary.BigEndian, uint16(len(encryptedPayload)))
	token.Write(encryptedPayload)
	return token.Bytes()
}
