package memmodelpub

import (
	"sync"
	"testing"
)

func TestLoadBeforePublish(t *testing.T) {
	var p Publisher
	if cfg, ok := p.Load(); ok || cfg != nil {
		t.Errorf("Load = %v, %v; want nil, false", cfg, ok)
	}
	if p.Ready() {
		t.Error("Ready = true before any Publish")
	}
}

func TestPublishThenLoad(t *testing.T) {
	var p Publisher
	BuildAndPublish(&p, "svc", 3, []string{"a", "b"})

	cfg, ok := p.Load()
	if !ok {
		t.Fatal("Load = false after Publish")
	}
	if cfg.Name != "svc" || cfg.Version != 3 || len(cfg.Tags) != 2 {
		t.Errorf("cfg = %+v", *cfg)
	}
	if !p.Ready() {
		t.Error("Ready = false after Publish")
	}
}

func TestPublishedConfigDoesNotAliasCaller(t *testing.T) {
	var p Publisher
	tags := []string{"a"}
	BuildAndPublish(&p, "svc", 1, tags)

	tags[0] = "MUTATED"
	cfg, _ := p.Load()
	if cfg.Tags[0] != "a" {
		t.Errorf("published config aliases the caller's slice: %q", cfg.Tags[0])
	}
}

func TestConcurrentPublishersAndReaders(t *testing.T) {
	var p Publisher
	var wg sync.WaitGroup

	for w := 0; w < 4; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				BuildAndPublish(&p, "svc", w*1000+i, []string{"t1", "t2", "t3"})
			}
		}(w)
	}

	for r := 0; r < 4; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 2000; i++ {
				cfg, ok := p.Load()
				if !ok {
					continue
				}
				// A partially published object would show up here as a
				// zero Name or a short Tags slice.
				if cfg.Name != "svc" {
					t.Errorf("observed partially initialised config: Name=%q", cfg.Name)
					return
				}
				if len(cfg.Tags) != 3 {
					t.Errorf("observed partially initialised config: Tags=%v", cfg.Tags)
					return
				}
			}
		}()
	}

	wg.Wait()
}

func TestLastPublishWins(t *testing.T) {
	var p Publisher
	BuildAndPublish(&p, "a", 1, nil)
	BuildAndPublish(&p, "b", 2, nil)

	cfg, _ := p.Load()
	if cfg.Name != "b" || cfg.Version != 2 {
		t.Errorf("cfg = %+v, want the second publish", *cfg)
	}
}

func TestIsLoader(t *testing.T) {
	var p Publisher
	var l Loader = &p
	if _, ok := l.Load(); ok {
		t.Error("Load = true on an empty publisher")
	}
}

func BenchmarkLoad(b *testing.B) {
	var p Publisher
	BuildAndPublish(&p, "svc", 1, []string{"a"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = p.Load()
		}
	})
}
