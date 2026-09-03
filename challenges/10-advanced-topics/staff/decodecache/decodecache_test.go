package decodecache

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type cfg struct {
	Host   string `env:"HOST"`
	Region string `env:"REGION"`
	Plain  string
	hidden string `env:"HIDDEN"`
	Count  int    `env:"COUNT"`
}

func TestDecode(t *testing.T) {
	var c cfg
	src := map[string]string{"HOST": "h", "REGION": "r", "HIDDEN": "x", "COUNT": "5"}
	if err := Decode(src, &c); err != nil {
		t.Fatal(err)
	}
	if c.Host != "h" || c.Region != "r" {
		t.Errorf("c = %+v, want Host=h Region=r", c)
	}
	if c.Plain != "" || c.hidden != "" || c.Count != 0 {
		t.Errorf("c = %+v: untagged, unexported and non-string fields must be left alone", c)
	}
}

func TestDecodeMissingKeys(t *testing.T) {
	c := cfg{Host: "keep"}
	if err := Decode(map[string]string{"REGION": "r"}, &c); err != nil {
		t.Fatal(err)
	}
	if c.Host != "keep" || c.Region != "r" {
		t.Errorf("c = %+v", c)
	}
}

func TestDecodeBadTarget(t *testing.T) {
	for _, dst := range []any{cfg{}, nil, (*cfg)(nil), new(int)} {
		if err := Decode(nil, dst); !errors.Is(err, ErrTarget) {
			t.Errorf("Decode(%#v) = %v, want ErrTarget", dst, err)
		}
	}
}

func TestDecodeConcurrent(t *testing.T) {
	const workers = 16
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make([]error, workers)
	got := make([]cfg, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			src := map[string]string{
				"HOST":   fmt.Sprintf("h%d", w),
				"REGION": fmt.Sprintf("r%d", w),
			}
			for i := 0; i < 200; i++ {
				var c cfg
				if err := Decode(src, &c); err != nil {
					errs[w] = err
					return
				}
				if c.Host != src["HOST"] || c.Region != src["REGION"] {
					errs[w] = fmt.Errorf("worker %d got %+v", w, c)
					return
				}
				got[w] = c
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

func TestDecodeManyTypes(t *testing.T) {
	type a struct {
		V string `env:"V"`
	}
	type b struct {
		W string `env:"W"`
	}
	var av a
	var bv b
	if err := Decode(map[string]string{"V": "1", "W": "2"}, &av); err != nil {
		t.Fatal(err)
	}
	if err := Decode(map[string]string{"V": "1", "W": "2"}, &bv); err != nil {
		t.Fatal(err)
	}
	if av.V != "1" || bv.W != "2" {
		t.Errorf("av = %+v, bv = %+v", av, bv)
	}
}
