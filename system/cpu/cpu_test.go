package cpu

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGet(c *C) {
	cpu, err := Get()
	c.Assert(err == nil, Equals, true)
	c.Assert(len(cpu) > 0, Equals, true)
}
