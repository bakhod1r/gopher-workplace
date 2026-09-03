// Package perworkerbuf — Gopher Workplace challenge.
package perworkerbuf

import (
	"strconv"
	"sync"
)

// RenderAll renders each row concurrently as comma-separated decimals
// and returns the results in input order.
//
// Each goroutine's scratch buffer must be a local that does not escape:
// one shared buffer is a race, and one heap buffer per row is garbage.
//
// Examples:
//
//	RenderAll([][]int{{1, 2}}) => []string{"1,2"}
func RenderAll(rows [][]int) []string {
	panic("not implemented")
}
