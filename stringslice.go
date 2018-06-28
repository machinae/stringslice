// Package stringslice provides utilities for working with string slices
package stringslice

// Unique returns a new slice with duplicate values silently discarded
func Unique(s []string) []string {
	u := make([]string, 0)
	m := make(map[string]bool)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
			u = append(u, v)
		}
	}
	return u
}

// Map returns the result of applying f to each element of the slice
func Map(s []string, f func(v string) string) []string {
	vs := make([]string, len(s))
	for i, v := range s {
		vs[i] = f(v)
	}
	return vs
}
