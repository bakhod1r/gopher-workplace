package schemaloader

import (
	"errors"
	"sync"
	"testing"
)

var errParse = errors.New("bad schema")

func TestLoaderLoad(t *testing.T) {
	cases := []struct {
		name    string
		calls   int
		schema  string
		err     error
		wantTry int
	}{
		{"never_queried", 0, "schema", nil, 0},
		{"first_query", 1, "schema", nil, 1},
		{"repeat_query", 3, "schema", nil, 1},
		{"failure_cached", 3, "", errParse, 1},
		{"many_queries", 20, "s", nil, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLoader(func() (string, error) { return tc.schema, tc.err })
			for i := 0; i < tc.calls; i++ {
				gotSchema, gotErr := l.Load()
				if gotSchema != tc.schema || !errors.Is(gotErr, tc.err) {
					t.Fatalf("Load() = %q, %v, want %q, %v", gotSchema, gotErr, tc.schema, tc.err)
				}
			}
			if got := l.Attempts(); got != tc.wantTry {
				t.Errorf("Attempts() = %d, want %d", got, tc.wantTry)
			}
		})
	}
}

func TestLoaderConcurrent(t *testing.T) {
	l := NewLoader(func() (string, error) { return "schema", nil })
	const queries = 32
	var wg sync.WaitGroup
	wg.Add(queries)
	for i := 0; i < queries; i++ {
		go func() {
			defer wg.Done()
			s, err := l.Load()
			if s != "schema" || err != nil {
				t.Errorf("Load() = %q, %v, want %q, nil", s, err, "schema")
			}
		}()
	}
	wg.Wait()
	if got := l.Attempts(); got != 1 {
		t.Errorf("Attempts() = %d, want 1", got)
	}
}
