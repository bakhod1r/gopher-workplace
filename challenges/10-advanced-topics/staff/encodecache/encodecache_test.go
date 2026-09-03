package encodecache

import (
	"bytes"
	"errors"
	"sync"
	"testing"
)

type user struct {
	Name   string
	Email  string
	Age    int
	hidden string
}

func TestEncode(t *testing.T) {
	got, err := Encode(nil, user{Name: "ann", Email: "a@b", Age: 3, hidden: "x"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, []byte("Name=ann&Email=a@b")) {
		t.Errorf("Encode = %q, want \"Name=ann&Email=a@b\"", got)
	}
}

func TestEncodeAppendsToDst(t *testing.T) {
	got, err := Encode([]byte("pre:"), user{Name: "x"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, []byte("pre:Name=x&Email=")) {
		t.Errorf("Encode = %q, want \"pre:Name=x&Email=\"", got)
	}
}

func TestEncodeNoStringFields(t *testing.T) {
	type nums struct{ A, B int }
	got, err := Encode(nil, nums{1, 2})
	if err != nil || len(got) != 0 {
		t.Errorf("Encode = %q, %v, want empty, nil", got, err)
	}
}

func TestEncodeBadKind(t *testing.T) {
	for _, v := range []any{nil, 3, []int{1}, &user{}} {
		if _, err := Encode(nil, v); !errors.Is(err, ErrKind) {
			t.Errorf("Encode(%#v) = %v, want ErrKind", v, err)
		}
	}
}

func TestEncodeConcurrent(t *testing.T) {
	const workers = 16
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make([]error, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			u := user{Name: string(rune('a' + w)), Email: "e"}
			want := []byte("Name=" + string(rune('a'+w)) + "&Email=e")
			buf := make([]byte, 0, 64)
			for i := 0; i < 200; i++ {
				got, err := Encode(buf[:0], u)
				if err != nil {
					errs[w] = err
					return
				}
				if !bytes.Equal(got, want) {
					errs[w] = errors.New("wrong output: " + string(got))
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

func TestEncodeSteadyStateAllocatesNothing(t *testing.T) {
	u := user{Name: "ann", Email: "a@b"}
	buf := make([]byte, 0, 128)
	Encode(buf[:0], u)
	var sink []byte
	n := testing.AllocsPerRun(200, func() { sink, _ = Encode(buf[:0], u) })
	_ = sink
	if n != 0 {
		t.Errorf("Encode made %v allocations, want 0: use the cached layout and the caller's buffer", n)
	}
}
