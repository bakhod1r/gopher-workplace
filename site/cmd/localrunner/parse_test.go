package main

import (
	"errors"
	"strings"
	"testing"
)

// events builds a `go test -json` stream from lines of {action, test, output}.
func events(lines ...string) []byte { return []byte(strings.Join(lines, "\n") + "\n") }

func TestParseTestJSONPass(t *testing.T) {
	out := events(
		`{"Action":"run","Test":"TestA"}`,
		`{"Action":"output","Test":"TestA","Output":"=== RUN   TestA\n"}`,
		`{"Action":"pass","Test":"TestA"}`,
		`{"Action":"run","Test":"TestB"}`,
		`{"Action":"pass","Test":"TestB"}`,
	)
	rep := parseTestJSON(out, "", nil)
	if !rep.OK || !rep.CompileOK {
		t.Fatalf("ok=%v compileOk=%v error=%q", rep.OK, rep.CompileOK, rep.Error)
	}
	if len(rep.Cases) != 2 {
		t.Fatalf("cases = %d, want 2", len(rep.Cases))
	}
	for _, c := range rep.Cases {
		if !c.Pass {
			t.Errorf("%s: pass = false", c.Name)
		}
	}
}

func TestParseTestJSONFailExtractsGotWant(t *testing.T) {
	out := events(
		`{"Action":"run","Test":"TestA"}`,
		`{"Action":"output","Test":"TestA","Output":"    a_test.go:12: Dedupe() = [1 1 2], want [1 2]\n"}`,
		`{"Action":"fail","Test":"TestA"}`,
	)
	rep := parseTestJSON(out, "", errors.New("exit status 1"))
	if rep.OK {
		t.Error("ok = true for a failing test")
	}
	if !rep.CompileOK {
		t.Error("compileOk = false; the package built fine")
	}
	if len(rep.Cases) != 1 {
		t.Fatalf("cases = %d, want 1", len(rep.Cases))
	}
	c := rep.Cases[0]
	if c.Got != "[1 1 2]" || c.Want != "[1 2]" {
		t.Errorf("got/want = %q/%q, want %q/%q", c.Got, c.Want, "[1 1 2]", "[1 2]")
	}
}

func TestParseTestJSONBuildFailure(t *testing.T) {
	out := events(
		`{"Action":"output","Output":"# example/p\n"}`,
		`{"Action":"output","Output":"./p.go:7:2: undefined: foo\n"}`,
		`{"Action":"output","Output":"FAIL\texample/p [build failed]\n"}`,
	)
	rep := parseTestJSON(out, "", errors.New("exit status 2"))
	if rep.CompileOK {
		t.Error("compileOk = true for a build failure")
	}
	if !strings.Contains(rep.Error, "undefined: foo") {
		t.Errorf("error = %q, want the compiler message", rep.Error)
	}
}

// A compile error can arrive as raw (non-JSON) text before test2json framing.
func TestParseTestJSONRawCompileError(t *testing.T) {
	rep := parseTestJSON([]byte("./p.go:7:2: undefined: foo\n"), "", errors.New("exit status 2"))
	if rep.CompileOK {
		t.Error("compileOk = true for a raw compile error")
	}
	if !strings.Contains(rep.Error, "undefined: foo") {
		t.Errorf("error = %q", rep.Error)
	}
}

func TestParseTestJSONSkipIsNotACase(t *testing.T) {
	out := events(
		`{"Action":"run","Test":"TestA"}`,
		`{"Action":"pass","Test":"TestA"}`,
		`{"Action":"run","Test":"TestSkipped"}`,
		`{"Action":"skip","Test":"TestSkipped"}`,
	)
	rep := parseTestJSON(out, "", nil)
	if len(rep.Cases) != 1 || rep.Cases[0].Name != "TestA" {
		t.Fatalf("cases = %+v, want only TestA", rep.Cases)
	}
	if !rep.OK {
		t.Error("ok = false; a skip must not fail the run")
	}
}

// WARN: lines are how guard tests flag hardcoding without failing; they must
// surface as warnings, which block a Submit from counting as solved.
func TestParseTestJSONCollectsWarnings(t *testing.T) {
	out := events(
		`{"Action":"run","Test":"TestGuard"}`,
		`{"Action":"output","Test":"TestGuard","Output":"    guard_test.go:9: WARN: looks hardcoded\n"}`,
		`{"Action":"pass","Test":"TestGuard"}`,
	)
	rep := parseTestJSON(out, "", nil)
	if len(rep.Warnings) != 1 || rep.Warnings[0] != "looks hardcoded" {
		t.Fatalf("warnings = %q", rep.Warnings)
	}
	if !rep.OK {
		t.Error("ok = false; a warning must not fail the run itself")
	}
}

func TestParseTestJSONNoTests(t *testing.T) {
	out := events(`{"Action":"output","Output":"testing: warning: no tests to run\n"}`)
	rep := parseTestJSON(out, "", nil)
	if rep.OK {
		t.Error("ok = true when nothing ran")
	}
	if rep.Error == "" {
		t.Error("error is empty; the package output should be surfaced")
	}
}

func TestExtractGotWantFallback(t *testing.T) {
	got, want := extractGotWant("=== RUN   TestA\n    a_test.go:3: boom\n--- FAIL: TestA\n")
	if got != "a_test.go:3: boom" || want != "" {
		t.Errorf("got/want = %q/%q", got, want)
	}
	if got, _ := extractGotWant(""); got != "test failed" {
		t.Errorf("empty output: got = %q, want %q", got, "test failed")
	}
	long := strings.Repeat("x", 300)
	g, _ := extractGotWant("    " + long + "\n")
	if len([]rune(g)) > 201 {
		t.Errorf("long line not condensed: %d runes", len([]rune(g)))
	}
}

func TestPassCount(t *testing.T) {
	r := report{Cases: []caseResult{{Pass: true}, {Pass: false}, {Pass: true}}}
	if got := passCount(r); got != 2 {
		t.Errorf("passCount = %d, want 2", got)
	}
}

// The idiom nearly every puzzle test uses: "Call(args) = X, want Y" with no
// literal "got". The UI's got/want columns depend on parsing it.
func TestExtractGotWantEqualsIdiom(t *testing.T) {
	cases := []struct {
		line, got, want string
	}{
		{"    a_test.go:12: Weekday(3) = \"Wed\", want \"Wednesday\"", `"Wed"`, `"Wednesday"`},
		{"    a_test.go:9: WholePart(3.7) = 4, want 3", "4", "3"},
		{"    a_test.go:4: len(Tags) = 2, want 0", "2", "0"},
		{"    a_test.go:7: input was mutated: got [1 1], want [1 2]", "[1 1]", "[1 2]"},
	}
	for _, c := range cases {
		got, want := extractGotWant(c.line + "\n")
		if got != c.got || want != c.want {
			t.Errorf("%q\n  got/want = %q/%q\n  wanted     %q/%q", c.line, got, want, c.got, c.want)
		}
	}
}
