package main

import (
	"encoding/json"
	"fmt"

	runner "github.com/gopher-workplace/site/runner"
)

const broken = `package dedupe
func Dedupe(in []int) []int {
	out := in
	seen := map[int]bool{}
	for _, v := range in {
		if seen[v] { continue }
		seen[v] = true
		out = append(out, v)
	}
	return out
}`

const fixed = `package dedupe
func Dedupe(in []int) []int {
	out := make([]int, 0, len(in))
	seen := make(map[int]bool, len(in))
	for _, v := range in {
		if seen[v] { continue }
		seen[v] = true
		out = append(out, v)
	}
	return out
}`

func main() {
	for _, tc := range []struct{ name, src string }{{"BROKEN", broken}, {"FIXED", fixed}} {
		rep := runner.RunDedupe(tc.src)
		b, _ := json.MarshalIndent(rep, "", "  ")
		fmt.Printf("=== %s ===\nOK=%v\n%s\n\n", tc.name, rep.OK, b)
	}
}
