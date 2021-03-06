# stringslice - Utilities for working with string slices

[![Go Report Card](https://goreportcard.com/badge/github.com/machinae/stringslice)](https://goreportcard.com/report/github.com/machinae/stringslice)

stringslice provides some Go utility functions for working with string slices.

All functions do not mutate the original slice, instead returning a new
slice.

## Documentation
Documentation is at [godoc](https://godoc.org/github.com/machinae/stringslice).

## Installation
```sh
go get github.com/machinae/stringslice
```

## Functions

### Unique
Returns a new slice with duplicate values discarded

### Map
Transforms the slice by passing it through a function

### Filter
Filters the slice, returning values for whicch a function returns true

### Compact
Removes empty strings

### Contains
Returns true if the slice contains the given string

### Copy
Creates a copy of the slice

### Diff
Returns values that are present in one of the slices but not both

### Intersect
Returns values that are present in both slices
