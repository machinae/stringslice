// Package stringslice provides utilities for working with string slices
package stringslice

import "strings"

// convert the slice to a set stored as a map
func toSet(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}
	return m
}

// convert the slice to a set indexed by lower case value
func toSetFold(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		k := strings.ToLower(v)
		if _, ok := m[k]; !ok {
			m[k] = true
		}
	}
	return m
}

// Returns values present in both sets
func setInter(s, s2 map[string]bool) map[string]bool {
	m := make(map[string]bool)
	for k := range s {
		if _, ok := s2[k]; ok {
			m[k] = true
		}
	}
	return m
}

// Returns values present in only one of the sets
func setDiff(s, s2 map[string]bool) map[string]bool {
	m := make(map[string]bool)
	for k := range s {
		if _, ok := s2[k]; !ok {
			m[k] = true
		}
	}
	for k := range s2 {
		if _, ok := s[k]; !ok {
			m[k] = true
		}
	}
	return m
}

// Returns string values in set
// Order of returned values is random
func fromSet(sm map[string]bool) []string {
	var vs []string
	for k := range sm {
		vs = append(vs, k)
	}
	return vs
}

// Unique returns a new slice with duplicate values silently discarded
func Unique(s []string) []string {
	return fromSet(toSet(s))
}

// Map returns the result of applying f to each element of the slice
func Map(s []string, f func(v string) string) []string {
	vs := make([]string, len(s))
	for i, v := range s {
		vs[i] = f(v)
	}
	return vs
}

// Filter returns a slice of values where f returns true
func Filter(s []string, f func(v string) bool) []string {
	vs := make([]string, 0)
	for _, v := range s {
		if f(v) {
			vs = append(vs, v)
		}
	}
	return vs
}

// Compact filters out empty strings
func Compact(s []string) []string {
	return Filter(s, func(v string) bool {
		return len(v) > 0
	})
}

// Contains returns true if the given string is present in the slice
func Contains(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

// ContainsFold is like Contains, but case-insensitive
func ContainsFold(s []string, val string) bool {
	for _, v := range s {
		if strings.EqualFold(v, val) {
			return true
		}
	}
	return false
}

// Copy creates a shallow copy of the slice
func Copy(s []string) []string {
	vs := make([]string, len(s))
	for i, v := range s {
		vs[i] = v
	}
	return vs
}

// Diff returns values that are present in one of the slices but not both
func Diff(s []string, s2 []string) []string {
	sm := setDiff(toSet(s), toSet(s2))
	return fromSet(sm)
}

// DiffFold is like Diff, but case insensitive
func DiffFold(s []string, s2 []string) []string {
	sm := setDiff(toSetFold(s), toSetFold(s2))
	return fromSet(sm)
}

// Intersect returns values that are present in both slices
func Intersect(s []string, s2 []string) []string {
	sm := setInter(toSet(s), toSet(s2))
	return fromSet(sm)
}

// Exclude returns a copy of the slice with values removed
func Exclude(s []string, vs ...string) []string {
	ss := toSet(vs)
	return Filter(s, func(str string) bool {
		return !ss[str]
	})
}

// Union combines unique values in both slices
func Union(s []string, s2 []string) []string {
	ss := toSet(s)
	for _, v := range s2 {
		ss[v] = true
	}
	return fromSet(ss)
}
