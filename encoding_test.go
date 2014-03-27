package otk_test

import (
	"bytes"
	"encoding/base64"
	"github.com/jkakar/otk"
	. "launchpad.net/gocheck"
)

var _ = Suite(&EncoderSuite{})

type EncoderSuite struct{}

// Encoder.Encode converts an OpenToken into a base64-encoded blob that can be
// transmitted over the wire.
func (s *EncoderSuite) TestEncodeAES128(c *C) {
	key, err := base64.StdEncoding.DecodeString("a66C9MvM8eY4qJKyCXKW+w==")
	c.Assert(err, IsNil)
	writer := new(bytes.Buffer)
	encoder := otk.NewEncoder(otk.AES128, key, writer)
	token := otk.OpenToken{"foo": []string{"bar"}, "bar": []string{"baz"}}
	err = encoder.Encode(token)
	c.Assert(err, IsNil)
	c.Assert(writer.String(), Equals, "UFRLAQK9THj0okLTUB663QrJFg5qA58IDhAb93ondvcx7sY6s44eszNqAAAga5W8Dc4XZwtsZ4qV3_lDI-Zn2_yadHHIhkGqNV5J9kw*")
}
