package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

/* #nosec */
func SortSlice[T constraints.Ordered](s []T) { // #nosec
	sort.Slice(s, func(i, j int) bool { // #nosec
		return s[i] < s[j] // #nosec
	}) // #nosec
} // #nosec
