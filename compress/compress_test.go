package xcompressions

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCompress(c *C) {
	data, err := Compress([]byte("Raed Shomali"))
	c.Assert(err, Equals, nil)

	decompressedData, err := Decompress(data)
	c.Assert(err, Equals, nil)

	c.Assert(string(decompressedData), Equals, "Raed Shomali")
}

func (s *MySuite) TestDecompress(c *C) {
	_, err := Decompress([]byte("Kaka"))
	c.Assert(err == nil, Equals, false)
}
