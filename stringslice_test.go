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

// Result of Filter with len(s) == 2
var filterTests = []sliceTest{
	{[]string{"aa", "bb"}, []string{"aa", "bb", "c"}},
	{[]string{}, []string{"a", "b", "c"}},
}

var compactTests = []sliceTest{
	{[]string{"aa", "bb"}, []string{"aa", "bb", "", ""}},
	{[]string{}, []string{"", "", ""}},
	{[]string{}, nil},
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

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	for _, ts := range filterTests {
		assert.EqualValues(ts.expected, Filter(ts.input, func(v string) bool {
			return len(v) == 2
		}))
	}
}

func TestCompact(t *testing.T) {
	assert := assert.New(t)
	for _, ts := range compactTests {
		assert.EqualValues(ts.expected, Compact(ts.input))
	}
}

func TestContains(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "b", "c"}
	assert.True(Contains(s, "b"))
	assert.False(Contains(s, "d"))
}
