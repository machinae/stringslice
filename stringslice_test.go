package stringslice

import (
	"sort"
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
		vs := Unique(ts.input)
		sort.Strings(vs)
		assert.EqualValues(ts.expected, vs)
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

func TestContainsFold(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "b", "c"}
	assert.True(ContainsFold(s, "b"))
	assert.True(ContainsFold(s, "B"))
	assert.False(ContainsFold(s, "d"))
	assert.False(ContainsFold(s, "D"))
}

func TestCopy(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "b", "c"}
	vs := Copy(s)
	s[2] = "d"
	assert.EqualValues([]string{"a", "b", "c"}, vs)
}

func TestDiff(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "b", "c", "d"}
	s2 := []string{"b", "c", "d", "e"}
	vs := Diff(s, s2)
	sort.Strings(vs)
	assert.EqualValues([]string{"a", "e"}, vs)
}

func TestDiffFold(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "B", "c", "d"}
	s2 := []string{"b", "c", "D", "e"}
	vs := DiffFold(s, s2)
	sort.Strings(vs)
	assert.EqualValues([]string{"a", "e"}, vs)
}

func TestIntersect(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "b", "c", "d"}
	s2 := []string{"b", "c", "d", "e"}
	vs := Intersect(s, s2)
	sort.Strings(vs)
	assert.EqualValues([]string{"b", "c", "d"}, vs)
}

func TestUnion(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "b", "c", "d"}
	s2 := []string{"b", "c", "d", "e"}
	vs := Union(s, s2)
	sort.Strings(vs)
	assert.EqualValues([]string{"a", "b", "c", "d", "e"}, vs)
}

func TestExclude(t *testing.T) {
	assert := assert.New(t)
	s := []string{"a", "b", "c", "d"}
	vs := Exclude(s, "d", "b")
	assert.EqualValues([]string{"a", "c"}, vs)
}
