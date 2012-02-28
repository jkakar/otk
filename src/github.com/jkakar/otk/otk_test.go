package otk_test

import (
	"github.com/jkakar/otk"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&OTKSuite{})

type OTKSuite struct{}

// OpenToken.Add creates a new key, if it isn't already present, and stores
// the specified value as []string value.
func (s *OTKSuite) TestAdd(c *C) {
	token := otk.OpenToken{}
	c.Assert(token["key"], IsNil)
	token.Add("key", "value")
	c.Assert(token["key"], Equals, []string{"value"})
}

// OpenToken.Add appends the specified value to the existing list of values,
// if the specified key already exists.
func (s *OTKSuite) TestAddToExistingKey(c *C) {
	token := otk.OpenToken{}
	token.Add("key", "value1")
	token.Add("key", "value2")
	c.Assert(token["key"], Equals, []string{"value1", "value2"})
}

// OpenToken.Del deletes all the values associated with the specified key.
func (s *OTKSuite) TestDel(c *C) {
	token := otk.OpenToken{"key": []string{"value"}}
	token.Del("key")
	c.Assert(token["key"], IsNil)
}

// OpenToken.Get returns the first value associated with the specified key.
func (s *OTKSuite) TestGet(c *C) {
	token := otk.OpenToken{"key": []string{"value1", "value2"}}
	c.Assert(token.Get("key"), Equals, "value1")
}

// OpenToken.Get returns "" if the specified key doesn't exist in the token.
func (s *OTKSuite) TestGetWithUnknownKey(c *C) {
	token := otk.OpenToken{}
	c.Assert(token.Get("key"), Equals, "")
}

// OpenToken.Set is identical to OpenToken.Add when the specified key doesn't
// exist in the token.
func (s *OTKSuite) TestSet(c *C) {
	token := otk.OpenToken{}
	c.Assert(token["key"], IsNil)
	token.Set("key", "value")
	c.Assert(token["key"], Equals, []string{"value"})
}

// OpenToken.Set replaces any existing values associated with the specified
// key.
func (s *OTKSuite) TestSetReplacesExistingValues(c *C) {
	token := otk.OpenToken{}
	token.Set("key", "value1")
	token.Set("key", "value2")
	c.Assert(token["key"], Equals, []string{"value2"})
}
