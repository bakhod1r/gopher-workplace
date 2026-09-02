package tabletestgen

import "testing"

func TestRunPasses(t *testing.T) {
	cases := []Case[int, int]{
		{Name: "zero", In: 0, Want: 0},
		{Name: "two", In: 2, Want: 4},
	}
	Run(t, "double", cases, func(n int) int { return n * 2 })
}

func TestRunCallsFnForEveryCase(t *testing.T) {
	calls := 0
	cases := []Case[int, int]{
		{Name: "a", In: 1, Want: 2},
		{Name: "b", In: 2, Want: 4},
		{Name: "c", In: 3, Want: 6},
	}
	Run(t, "double", cases, func(n int) int { calls++; return n * 2 })
	if calls != len(cases) {
		t.Errorf("fn called %d times, want %d", calls, len(cases))
	}
}

func TestRunEmptyTable(t *testing.T) {
	Run(t, "none", []Case[string, string]{}, func(s string) string { return s })
}

func TestRunStrings(t *testing.T) {
	cases := []Case[string, int]{{Name: "len", In: "abc", Want: 3}}
	Run(t, "length", cases, func(s string) int { return len(s) })
}
