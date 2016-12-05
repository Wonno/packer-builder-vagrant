package tabix

import (
	"bytes"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(&S{})

// Generated by `echo $'chr1\t1\t100' | bgzip -c > t.bed.gz; tabix t.bed.gz'.
// See https://github.com/biogo/hts/issues/19.
var issue19TestData = []byte{
	0x54, 0x42, 0x49, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, // |TBI.............|
	0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |........#.......|
	0x05, 0x00, 0x00, 0x00, 0x63, 0x68, 0x72, 0x31, 0x00, 0x02, 0x00, 0x00, 0x00, 0x49, 0x12, 0x00, // |....chr1.....I..|
	0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, // |...............%|
	0x00, 0x00, 0x00, 0x00, 0x00, 0x4a, 0x92, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |.....J..........|
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, // |.......%........|
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, // |................|
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |................|
	0x00, /*                                                                                     */ // |.|
}

func (s *S) TestIssue19(c *check.C) {
	idx, err := ReadFrom(bytes.NewReader(issue19TestData))
	c.Assert(err, check.Equals, nil)

	chunks, err := idx.Chunks("chr1", 1, 19999999)
	c.Assert(err, check.Equals, nil)
	c.Check(len(chunks), check.Not(check.Equals), 0)
}
