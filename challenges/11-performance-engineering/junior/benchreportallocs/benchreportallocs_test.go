package benchreportallocs

import "testing"

func TestReport(t *testing.T) {
	cases := []struct {
		bytes, allocs uint64
		iters         int
		want          string
	}{
		{2048, 8, 4, "512 B/op\t2 allocs/op"},
		{0, 0, 10, "0 B/op\t0 allocs/op"},
		{10, 3, 4, "2 B/op\t0 allocs/op"},
		{1, 1, 1, "1 B/op\t1 allocs/op"},
	}
	for _, c := range cases {
		if got := Report(c.bytes, c.allocs, c.iters); got != c.want {
			t.Errorf("Report(%d, %d, %d) = %q, want %q", c.bytes, c.allocs, c.iters, got, c.want)
		}
	}
}

func TestReportNonPositiveIters(t *testing.T) {
	for _, n := range []int{0, -1} {
		if got := Report(999, 999, n); got != "0 B/op\t0 allocs/op" {
			t.Errorf("Report(_, _, %d) = %q, want zeros", n, got)
		}
	}
}
