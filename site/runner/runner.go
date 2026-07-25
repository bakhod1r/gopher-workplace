// Package poc runs a slice-challenge solution entirely inside a Go interpreter
// (yaegi), so it can execute in the browser via WebAssembly with no server.
//
// This POC targets the "dedupe" junior puzzle: interpret the user's source,
// then call Dedupe(...) for each test case and compare against the expected
// result. Pure-Go, no cgo, no real `go test` — enough to prove in-browser run.
package runner

import (
	"fmt"
	"go/format"
	"reflect"
	"strconv"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// CaseResult is one test-case outcome, JSON-friendly for the JS bridge.
type CaseResult struct {
	Name string `json:"name"`
	Pass bool   `json:"pass"`
	Got  string `json:"got"`
	Want string `json:"want"`
}

// Report is the full run outcome.
type Report struct {
	OK        bool         `json:"ok"`
	CompileOK bool         `json:"compileOk"`
	Error     string       `json:"error"`
	Cases     []CaseResult `json:"cases"`
}

type dedupeCase struct {
	name string
	in   []int
	want []int
}

var dedupeCases = []dedupeCase{
	{"empty", []int{}, []int{}},
	{"no dups", []int{1, 2, 3}, []int{1, 2, 3}},
	{"adjacent dups", []int{1, 1, 2, 3, 3, 3}, []int{1, 2, 3}},
	{"non-adjacent", []int{5, 4, 5, 4}, []int{5, 4}},
	{"all same", []int{7, 7, 7}, []int{7}},
	{"order preserved", []int{3, 1, 2, 1, 3}, []int{3, 1, 2}},
}

// RunDedupe interprets userSrc (which must define Dedupe) and checks it.
func RunDedupe(userSrc string) Report {
	src := forceMainPackage(userSrc)

	i := interp.New(interp.Options{})
	if err := i.Use(stdlib.Symbols); err != nil {
		return Report{Error: "interp init: " + err.Error()}
	}

	if _, err := i.Eval(src); err != nil {
		return Report{CompileOK: false, Error: err.Error()}
	}

	rep := Report{CompileOK: true, OK: true}
	for _, c := range dedupeCases {
		expr := fmt.Sprintf("Dedupe(%#v)", c.in)
		v, err := i.Eval(expr)
		if err != nil {
			rep.OK = false
			rep.Cases = append(rep.Cases, CaseResult{Name: c.name, Pass: false, Got: "error: " + err.Error(), Want: fmt.Sprintf("%v", c.want)})
			continue
		}
		got := toIntSlice(v)
		pass := equalIgnoringEmpty(got, c.want)
		if !pass {
			rep.OK = false
		}
		rep.Cases = append(rep.Cases, CaseResult{
			Name: c.name, Pass: pass,
			Got:  fmt.Sprintf("%v", got),
			Want: fmt.Sprintf("%v", c.want),
		})
	}
	return rep
}

// CustomResult is the outcome of running Dedupe on a user-supplied input.
type CustomResult struct {
	CompileOK bool   `json:"compileOk"`
	OK        bool   `json:"ok"`
	Error     string `json:"error"`
	Input     string `json:"input"`
	Output    string `json:"output"`
}

// RunDedupeCustom interprets userSrc and runs Dedupe on a single caller-supplied
// input (a JSON int array like "[1,2,2,3]"), returning the produced output.
// No expected value — LeetCode "custom input" style.
func RunDedupeCustom(userSrc, inputJSON string) CustomResult {
	in, err := parseIntSlice(inputJSON)
	if err != nil {
		return CustomResult{Error: "bad input: " + err.Error()}
	}

	i := interp.New(interp.Options{})
	if err := i.Use(stdlib.Symbols); err != nil {
		return CustomResult{Error: "interp init: " + err.Error()}
	}
	if _, err := i.Eval(forceMainPackage(userSrc)); err != nil {
		return CustomResult{CompileOK: false, Error: err.Error()}
	}

	v, err := i.Eval(fmt.Sprintf("Dedupe(%#v)", in))
	if err != nil {
		return CustomResult{CompileOK: true, Error: err.Error(), Input: fmt.Sprintf("%v", in)}
	}
	return CustomResult{
		CompileOK: true, OK: true,
		Input:  fmt.Sprintf("%v", in),
		Output: fmt.Sprintf("%v", toIntSlice(v)),
	}
}

// ---- plan-limits puzzle (variables & constants: iota, shadowing) ----

type planCase struct {
	name string
	expr string
	want int
}

var planCases = []planCase{
	{"tiers distinct & ascending", "int(Free)*100 + int(Pro)*10 + int(Enterprise)", 12}, // 0,1,2 -> 012
	{"free", "Limit(Free)", 60},
	{"pro", "Limit(Pro)", 600},
	{"enterprise", "Limit(Enterprise)", 6000},
	{"unknown => free fallback", "Limit(Tier(99))", 60},
}

// RunPlanLimits interprets userSrc (which must define Tier, Free/Pro/Enterprise
// and Limit) and checks each case.
func RunPlanLimits(userSrc string) Report {
	src := forceMainPackage(userSrc)

	i := interp.New(interp.Options{})
	if err := i.Use(stdlib.Symbols); err != nil {
		return Report{Error: "interp init: " + err.Error()}
	}
	if _, err := i.Eval(src); err != nil {
		return Report{CompileOK: false, Error: err.Error()}
	}

	rep := Report{CompileOK: true, OK: true}
	for _, c := range planCases {
		v, err := i.Eval(c.expr)
		if err != nil {
			rep.OK = false
			rep.Cases = append(rep.Cases, CaseResult{Name: c.name, Pass: false,
				Got: "error: " + err.Error(), Want: strconv.Itoa(c.want)})
			continue
		}
		got := int(v.Int())
		pass := got == c.want
		if !pass {
			rep.OK = false
		}
		rep.Cases = append(rep.Cases, CaseResult{Name: c.name, Pass: pass,
			Got: strconv.Itoa(got), Want: strconv.Itoa(c.want)})
	}
	return rep
}

// RunPlanLimitsCustom runs Limit on a single caller-supplied tier value
// (an integer like "1", or "99" for the unknown case).
func RunPlanLimitsCustom(userSrc, input string) CustomResult {
	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return CustomResult{Error: "bad input: want an integer tier, got " + strconv.Quote(input)}
	}
	i := interp.New(interp.Options{})
	if err := i.Use(stdlib.Symbols); err != nil {
		return CustomResult{Error: "interp init: " + err.Error()}
	}
	if _, err := i.Eval(forceMainPackage(userSrc)); err != nil {
		return CustomResult{CompileOK: false, Error: err.Error()}
	}
	v, err := i.Eval(fmt.Sprintf("Limit(Tier(%d))", n))
	if err != nil {
		return CustomResult{CompileOK: true, Error: err.Error(), Input: fmt.Sprintf("Tier(%d)", n)}
	}
	return CustomResult{CompileOK: true, OK: true,
		Input: fmt.Sprintf("Tier(%d)", n), Output: strconv.Itoa(int(v.Int()))}
}

// parseIntSlice accepts "[1, 2, 3]" or "1,2,3" (whitespace tolerant).
func parseIntSlice(s string) ([]int, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	s = strings.TrimSpace(s)
	if s == "" {
		return []int{}, nil
	}
	parts := strings.Split(s, ",")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("%q is not an integer", p)
		}
		out = append(out, n)
	}
	return out, nil
}

// forceMainPackage rewrites the package clause to main so yaegi evaluates the
// top-level declarations into a usable scope.
func forceMainPackage(src string) string {
	lines := strings.Split(src, "\n")
	for idx, ln := range lines {
		if strings.HasPrefix(strings.TrimSpace(ln), "package ") {
			lines[idx] = "package main"
			break
		}
	}
	return strings.Join(lines, "\n")
}

func toIntSlice(v reflect.Value) []int {
	if !v.IsValid() || v.Kind() != reflect.Slice {
		return nil
	}
	out := make([]int, v.Len())
	for i := 0; i < v.Len(); i++ {
		out[i] = int(v.Index(i).Int())
	}
	return out
}

// equalIgnoringEmpty treats nil and empty slices as equal (both len 0).
func equalIgnoringEmpty(got, want []int) bool {
	if len(got) == 0 && len(want) == 0 {
		return true
	}
	return reflect.DeepEqual(got, want)
}

// FormatResult is the outcome of gofmt-ing user source.
type FormatResult struct {
	OK     bool   `json:"ok"`
	Source string `json:"source"`
	Error  string `json:"error"`
}

// Format runs the real gofmt (go/format) over userSrc. On a parse error the
// original source is returned unchanged with OK=false and the compiler error.
func Format(userSrc string) FormatResult {
	out, err := format.Source([]byte(userSrc))
	if err != nil {
		return FormatResult{OK: false, Source: userSrc, Error: err.Error()}
	}
	return FormatResult{OK: true, Source: string(out)}
}
