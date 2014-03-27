// Package otk provides support for OpenToken (OTK).  OTK is a format for the
// lightweight, secure, cross-application exchange of key-value pairs between
// applications that use HTTP as the transport protocol.  The format is
// designed primarily for use as an HTTP cookie or query parameter, but can
// also be used in other scenarios that require a compact, application-neutral
// token.
//
// This implementation follows the recommendations from the draft
// specification:
//
// http://tools.ietf.org/html/draft-smith-opentoken-02
package otk

// A representation of the key/value pairs in an OpenToken.
type OpenToken map[string][]string

// Add the key/value pair to the token.  It appends to any existing values
// already associated with key.
func (token OpenToken) Add(key, value string) {
	token[key] = append(token[key], value)
}

// Delete the values associated with key.
func (token OpenToken) Del(key string) {
	delete(token, key)
}

// Get the first value associated with the given key.  If there are no values
// associated with the key, Get returns "".  To access multiple values of a
// key, access the map directly.
func (token OpenToken) Get(key string) string {
	if values := token[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

// Set the entries associated with key to the single element value.  It
// replaces any existing values associated with key.
func (token OpenToken) Set(key, value string) {
	token[key] = []string{value}
}
