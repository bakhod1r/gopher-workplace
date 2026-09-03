package typeregistry

import (
	"errors"
	"sync"
	"testing"
)

type job struct {
	ID   int
	Name string
}

type task struct{ N int }

func init() {
	Register("job", job{})
	Register("task", task{})
	Register("int", 0)
}

func TestNewReturnsAPointerToTheRegisteredType(t *testing.T) {
	v, err := New("job")
	if err != nil {
		t.Fatal(err)
	}
	p, ok := v.(*job)
	if !ok {
		t.Fatalf("New = %T, want *job", v)
	}
	if *p != (job{}) {
		t.Errorf("New = %+v, want the zero job", *p)
	}
	p.ID = 7
}

func TestNewGivesDistinctValues(t *testing.T) {
	a, _ := New("job")
	b, _ := New("job")
	if a.(*job) == b.(*job) {
		t.Error("New returned the same pointer twice")
	}
	a.(*job).ID = 1
	if b.(*job).ID != 0 {
		t.Error("the two values share storage")
	}
}

func TestNewOtherTypes(t *testing.T) {
	v, err := New("task")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := v.(*task); !ok {
		t.Errorf("New = %T, want *task", v)
	}
	v, err = New("int")
	if err != nil {
		t.Fatal(err)
	}
	if p, ok := v.(*int); !ok || *p != 0 {
		t.Errorf("New = %T, want *int holding 0", v)
	}
}

func TestNewUnknown(t *testing.T) {
	if _, err := New("nope"); !errors.Is(err, ErrUnknown) {
		t.Errorf("err = %v, want ErrUnknown", err)
	}
	if _, err := New(""); !errors.Is(err, ErrUnknown) {
		t.Errorf("err = %v, want ErrUnknown", err)
	}
}

func TestNewConcurrent(t *testing.T) {
	const workers = 16
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make([]error, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				v, err := New("job")
				if err != nil {
					errs[w] = err
					return
				}
				p := v.(*job)
				p.ID = w
				if p.ID != w {
					errs[w] = errors.New("value is shared between goroutines")
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
