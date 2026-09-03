package boundedcache

import (
	"bytes"
	"sync"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	c := NewCache(4)
	c.Put("a", []byte("one"))
	got, ok := c.Get("a")
	if !ok || !bytes.Equal(got, []byte("one")) {
		t.Errorf("Get = %q, %v, want \"one\", true", got, ok)
	}
	if _, ok := c.Get("missing"); ok {
		t.Error("Get(missing) reported ok, want false")
	}
}

func TestPutCopiesTheValue(t *testing.T) {
	c := NewCache(4)
	buf := []byte("first")
	c.Put("k", buf)
	copy(buf, "SECON")
	got, _ := c.Get("k")
	if !bytes.Equal(got, []byte("first")) {
		t.Errorf("Get = %q, want \"first\": the cache stored the caller's buffer", got)
	}
}

func TestEvictsOldest(t *testing.T) {
	c := NewCache(2)
	c.Put("a", []byte("1"))
	c.Put("b", []byte("2"))
	c.Put("c", []byte("3"))
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
	if _, ok := c.Get("a"); ok {
		t.Error("the oldest entry was not evicted")
	}
	for _, k := range []string{"b", "c"} {
		if _, ok := c.Get(k); !ok {
			t.Errorf("%q was evicted, want it kept", k)
		}
	}
}

func TestOverwriteDoesNotEvict(t *testing.T) {
	c := NewCache(2)
	c.Put("a", []byte("1"))
	c.Put("b", []byte("2"))
	c.Put("a", []byte("updated"))
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
	got, ok := c.Get("a")
	if !ok || !bytes.Equal(got, []byte("updated")) {
		t.Errorf("Get(a) = %q, %v, want \"updated\", true", got, ok)
	}
	if _, ok := c.Get("b"); !ok {
		t.Error("b was evicted by an overwrite")
	}
}

func TestStaysBoundedUnderLoad(t *testing.T) {
	c := NewCache(8)
	for i := 0; i < 5000; i++ {
		c.Put(string(rune('a'+i%26))+string(rune('a'+i/26%26)), []byte("payload"))
		if c.Len() > 8 {
			t.Fatalf("Len = %d after %d puts, want at most 8", c.Len(), i+1)
		}
	}
}

func TestConcurrentPuts(t *testing.T) {
	c := NewCache(16)
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			buf := make([]byte, 8)
			for i := 0; i < 500; i++ {
				for j := range buf {
					buf[j] = byte('a' + w)
				}
				c.Put(string(rune('a'+w))+string(rune('0'+i%10)), buf)
				c.Get(string(rune('a' + w)))
			}
		}(w)
	}
	wg.Wait()
	if c.Len() > 16 {
		t.Errorf("Len = %d, want at most 16", c.Len())
	}
}
