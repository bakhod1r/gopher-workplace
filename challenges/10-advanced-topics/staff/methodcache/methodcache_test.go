package methodcache

import (
	"errors"
	"sync"
	"testing"
)

type adder struct{ Base int }

func (a adder) Add(n int) int         { return a.Base + n }
func (a adder) Twice(n int) int       { return 2 * n }
func (a adder) Name() string          { return "adder" }
func (a adder) Pair(n int) (int, int) { return n, n }

type ptrAdder struct{ Base int }

func (p *ptrAdder) Add(n int) int { return p.Base + n }

func TestCallNamed(t *testing.T) {
	if got, err := CallNamed(adder{Base: 2}, "Add", 3); err != nil || got != 5 {
		t.Errorf("CallNamed = %d, %v, want 5, nil", got, err)
	}
	if got, err := CallNamed(adder{}, "Twice", 4); err != nil || got != 8 {
		t.Errorf("CallNamed = %d, %v, want 8, nil", got, err)
	}
}

func TestCallNamedPointerReceiver(t *testing.T) {
	if got, err := CallNamed(&ptrAdder{Base: 10}, "Add", 5); err != nil || got != 15 {
		t.Errorf("CallNamed = %d, %v, want 15, nil", got, err)
	}
	if _, err := CallNamed(ptrAdder{Base: 10}, "Add", 5); !errors.Is(err, ErrMethod) {
		t.Error("a pointer-receiver method must not be reachable on the value")
	}
}

func TestCallNamedBadShapes(t *testing.T) {
	cases := []struct {
		name   string
		v      any
		method string
	}{
		{"missing", adder{}, "Nope"},
		{"wrong results", adder{}, "Name"},
		{"too many results", adder{}, "Pair"},
		{"nil value", nil, "Add"},
		{"no methods", 3, "Add"},
	}
	for _, c := range cases {
		if _, err := CallNamed(c.v, c.method, 1); !errors.Is(err, ErrMethod) {
			t.Errorf("%s: err = %v, want ErrMethod", c.name, err)
		}
	}
}

func TestCallNamedConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	errs := make([]error, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				got, err := CallNamed(adder{Base: w}, "Add", i)
				if err != nil {
					errs[w] = err
					return
				}
				if got != w+i {
					errs[w] = errors.New("wrong result")
					return
				}
			}
		}(w)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCallNamedRepeatsUseTheCache(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if got, err := CallNamed(adder{Base: 1}, "Add", i); err != nil || got != 1+i {
			t.Fatalf("iteration %d: %d, %v", i, got, err)
		}
	}
}
