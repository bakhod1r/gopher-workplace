package dbpoolinit

import (
	"sync"
	"testing"
)

func TestProviderPool(t *testing.T) {
	cases := []struct {
		name  string
		calls int
		dsn   string
		want  int
	}{
		{"never_queried", 0, "db", 0},
		{"first_query", 1, "db", 1},
		{"two_queries", 2, "postgres://x", 1},
		{"ten_queries", 10, "mysql://y", 1},
		{"hundred_queries", 100, "sqlite://z", 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewProvider(func() *Pool { return &Pool{DSN: tc.dsn} })
			var first *Pool
			for i := 0; i < tc.calls; i++ {
				got := p.Pool()
				if got.DSN != tc.dsn {
					t.Fatalf("Pool().DSN = %q, want %q", got.DSN, tc.dsn)
				}
				if first == nil {
					first = got
				} else if got != first {
					t.Fatal("Pool() returned a different pool on a later call")
				}
			}
			if got := p.Opens(); got != tc.want {
				t.Errorf("Opens() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestProviderConcurrent(t *testing.T) {
	p := NewProvider(func() *Pool { return &Pool{DSN: "db"} })
	const handlers = 32
	pools := make(chan *Pool, handlers)
	var wg sync.WaitGroup
	wg.Add(handlers)
	for i := 0; i < handlers; i++ {
		go func() {
			defer wg.Done()
			pools <- p.Pool()
		}()
	}
	wg.Wait()
	close(pools)

	first := <-pools
	for got := range pools {
		if got != first {
			t.Fatal("concurrent callers received different pools")
		}
	}
	if got := p.Opens(); got != 1 {
		t.Errorf("Opens() = %d, want 1", got)
	}
}
