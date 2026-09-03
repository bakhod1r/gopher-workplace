// Package sliceheadercopy — Gopher Workplace challenge.
package sliceheadercopy

// Fill overwrites every element of s with v. The slice header is copied by
// value, but it points at the caller's array, so the writes are visible to
// the caller and nothing is allocated.
//
// Examples:
//
//	Fill(s, 7) sets every element of the caller's s to 7
func Fill(s []int, v int) {
	panic("not implemented")
}

// AppendLocal appends v to s without returning anything. Because the header
// is a copy, the caller's length never changes — this function is deliberately
// useless, and the test pins that behaviour down.
//
// Examples:
//
//	AppendLocal(s, 7) leaves len(s) unchanged for the caller
func AppendLocal(s []int, v int) {
	panic("not implemented")
}
