package net

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGet(c *C) {
	net, err := Get()
	c.Assert(err == nil, Equals, true)
	c.Assert(len(net) > 0, Equals, true)
	fmt.Println(net)
}
