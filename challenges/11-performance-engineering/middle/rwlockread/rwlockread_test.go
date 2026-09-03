package rwlockread

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetAndReplace(t *testing.T) {
	var c Config
	if _, ok := c.Get("k"); ok {
		t.Error("Get on an empty config reported found")
	}
	if got := c.Version(); got != 0 {
		t.Errorf("Version = %d, want 0", got)
	}
	c.Replace(map[string]string{"k": "v"})
	if v, ok := c.Get("k"); v != "v" || !ok {
		t.Errorf("Get = %q, %v, want v, true", v, ok)
	}
	if got := c.Version(); got != 1 {
		t.Errorf("Version = %d, want 1", got)
	}
}

func TestReplaceCopiesTheInput(t *testing.T) {
	var c Config
	in := map[string]string{"k": "v"}
	c.Replace(in)
	in["k"] = "mutated"
	in["extra"] = "x"
	if v, _ := c.Get("k"); v != "v" {
		t.Errorf("Get = %q, want v — Replace must copy the map", v)
	}
	if _, ok := c.Get("extra"); ok {
		t.Error("the live config saw a key added to the caller's map")
	}
}

func TestSnapshotIsAConsistentPair(t *testing.T) {
	var c Config
	c.Replace(map[string]string{"a": "1"})
	vals, ver := c.Snapshot()
	c.Replace(map[string]string{"b": "2"})
	if ver != 1 {
		t.Errorf("version = %d, want 1", ver)
	}
	if _, ok := vals["a"]; !ok {
		t.Error("the snapshot changed after a later Replace")
	}
	if _, ok := vals["b"]; ok {
		t.Error("the snapshot picked up a later write")
	}
}

func TestSnapshotOfEmptyConfig(t *testing.T) {
	var c Config
	vals, ver := c.Snapshot()
	if vals == nil || len(vals) != 0 || ver != 0 {
		t.Errorf("Snapshot = %v, %d, want an empty non-nil map and 0", vals, ver)
	}
}

func TestConcurrentReadersAndOneWriter(t *testing.T) {
	var c Config
	c.Replace(map[string]string{"k": "v0"})
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				if _, ok := c.Get("k"); !ok {
					t.Error("key disappeared during a Replace")
					return
				}
				_ = c.Version()
				_, _ = c.Snapshot()
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			c.Replace(map[string]string{"k": fmt.Sprintf("v%d", i)})
		}
	}()
	wg.Wait()
	if got := c.Version(); got != 101 {
		t.Errorf("Version = %d, want 101", got)
	}
}
