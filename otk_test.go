package otk_test

import (
	"github.com/jkakar/otk"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&OpenTokenSuite{})

type OpenTokenSuite struct{}

// OpenToken.Add creates a new key, if it isn't already present, and stores
// the specified value as []string value.
func (s *OpenTokenSuite) TestAdd(c *C) {
	token := otk.OpenToken{}
	c.Assert(token["key"], IsNil)
	token.Add("key", "value")
	c.Assert(token["key"], Equals, []string{"value"})
}

// OpenToken.Add appends the specified value to the existing list of values,
// if the specified key already exists.
func (s *OpenTokenSuite) TestAddToExistingKey(c *C) {
	token := otk.OpenToken{}
	token.Add("key", "value1")
	token.Add("key", "value2")
	c.Assert(token["key"], Equals, []string{"value1", "value2"})
}

// OpenToken.Del deletes all the values associated with the specified key.
func (s *OpenTokenSuite) TestDel(c *C) {
	token := otk.OpenToken{"key": []string{"value"}}
	token.Del("key")
	c.Assert(token["key"], IsNil)
}

// OpenToken.Get returns the first value associated with the specified key.
func (s *OpenTokenSuite) TestGet(c *C) {
	token := otk.OpenToken{"key": []string{"value1", "value2"}}
	c.Assert(token.Get("key"), Equals, "value1")
}

// OpenToken.Get returns "" if the specified key doesn't exist in the token.
func (s *OpenTokenSuite) TestGetWithUnknownKey(c *C) {
	token := otk.OpenToken{}
	c.Assert(token.Get("key"), Equals, "")
}

// OpenToken.Set is identical to OpenToken.Add when the specified key doesn't
// exist in the token.
func (s *OpenTokenSuite) TestSet(c *C) {
	token := otk.OpenToken{}
	c.Assert(token["key"], IsNil)
	token.Set("key", "value")
	c.Assert(token["key"], Equals, []string{"value"})
}

// OpenToken.Set replaces any existing values associated with the specified
// key.
func (s *OpenTokenSuite) TestSetReplacesExistingValues(c *C) {
	token := otk.OpenToken{}
	token.Set("key", "value1")
	token.Set("key", "value2")
	c.Assert(token["key"], Equals, []string{"value2"})
}
