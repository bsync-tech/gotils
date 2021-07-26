package ip

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetPublicIP(c *C) {
	ip, err := GetPublicIP(nil)
	c.Assert(err == nil, Equals, true)
	c.Assert(ip != "", Equals, true)
}
