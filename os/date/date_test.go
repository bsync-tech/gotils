package date

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestParse(c *C) {
	_, err := Parse("2021/12/13 14:45:00")
	c.Assert(err == nil, Equals, true)
}
