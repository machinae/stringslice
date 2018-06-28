package stringslice

import (
	"strings"
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

// result of Map with strings.ToUpper
var mapTests = []sliceTest{
	{[]string{"A", "B", "C"}, []string{"a", "b", "c"}},
	{[]string{"A", "B", ""}, []string{"a", "B", ""}},
}

func TestUnique(t *testing.T) {
	assert := assert.New(t)
	for _, ts := range uniqueTests {
		assert.EqualValues(ts.expected, Unique(ts.input))
	}
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	for _, ts := range mapTests {
		assert.EqualValues(ts.expected, Map(ts.input, strings.ToUpper))
	}
}
