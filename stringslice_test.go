package stringslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type sliceTest struct {
	expected []string
	input    []string
}

var uniqueTests = []sliceTest{
	{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "c"}, []string{"a", "b", "c", "b"}},
}

func TestUnique(t *testing.T) {
	assert := assert.New(t)
	for _, ts := range uniqueTests {
		assert.EqualValues(ts.expected, Unique(ts.input))
	}
}
