package disk

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
	disk, err := Get()
	c.Assert(err == nil, Equals, true)
	c.Assert(len(disk) > 0, Equals, true)
}
