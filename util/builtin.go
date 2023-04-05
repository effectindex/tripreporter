// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package util

// SliceEqual returns true if all bytes of a and b are the same
func SliceEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// SliceRemoveIndex will remove index i from s
func SliceRemoveIndex[T comparable](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
