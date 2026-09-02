package iowriter

import (
	"bytes"
	"errors"
	"testing"
)

type failWriter struct {
	after int
	n     int
}

var errBoom = errors.New("boom")

func (f *failWriter) Write(p []byte) (int, error) {
	if f.n >= f.after {
		return 0, errBoom
	}
	f.n++
	return len(p), nil
}

func TestWriteReport(t *testing.T) {
	cases := []struct {
		name  string
		title string
		items []string
		want  string
	}{
		{"one_item", "T", []string{"a"}, "T\n- a\n"},
		{"no_items", "T", nil, "T\n"},
		{"two_items", "Report", []string{"x", "y"}, "Report\n- x\n- y\n"},
		{"empty_title", "", []string{"a"}, "\n- a\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			n, err := WriteReport(&buf, tc.title, tc.items)
			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}
			if buf.String() != tc.want {
				t.Errorf("output = %q, want %q", buf.String(), tc.want)
			}
			if n != len(tc.want) {
				t.Errorf("n = %d, want %d", n, len(tc.want))
			}
		})
	}
}

func TestWriteReportError(t *testing.T) {
	w := &failWriter{after: 1}
	n, err := WriteReport(w, "T", []string{"a", "b"})
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
	if n != 2 {
		t.Errorf("n = %d, want 2 (bytes written before the failure)", n)
	}
}
