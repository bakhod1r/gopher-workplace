// Package builderreuse — Gopher Workplace challenge.
package builderreuse

import (
	"strconv"
	"strings"
)

// RenderLines renders each row as its values joined by '-'.
//
// The builder is per-call state: reset it between rows instead of
// constructing one per row, and reserve its capacity once.
//
// Examples:
//
//	RenderLines([][]int{{1, 2}}) => []string{"1-2"}
func RenderLines(rows [][]int) []string {
	panic("not implemented")
}
