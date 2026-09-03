// Package foldjoinbug — Gopher Workplace challenge.
package foldjoinbug

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

// Sample is one profile sample: the call stack, caller first, and its value.
type Sample struct {
	Stack []string
	Value int64
}

// Fold renders the samples in the collapsed-stack format flame graph tools
// read: one line per distinct stack, "frame;frame;frame value", identical
// stacks summed. Two stacks with the same frames in a different order are
// different call paths and must stay apart.
//
// Examples:
//
//	Fold([{["a","b"], 3}]) => ["a;b 3"]
func Fold(samples []Sample) []string {
	totals := make(map[string]int64)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) == 0 {
			continue
		}
		// CHANGE CODE BELOW THIS LINE
		frames := slices.Sorted(slices.Values(s.Stack))
		// CHANGE CODE ABOVE THIS LINE
		totals[strings.Join(frames, ";")] += s.Value
	}
	type row struct {
		key string
		val int64
	}
	rows := make([]row, 0, len(totals))
	for k, v := range totals {
		rows = append(rows, row{k, v})
	}
	slices.SortFunc(rows, func(a, b row) int {
		if c := cmp.Compare(b.val, a.val); c != 0 {
			return c
		}
		return cmp.Compare(a.key, b.key)
	})
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s %d", r.key, r.val))
	}
	return out
}
