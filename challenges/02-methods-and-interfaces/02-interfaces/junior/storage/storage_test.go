package storage

import "testing"

func TestPutGet(t *testing.T) {
	s := NewMemStore()
	if _, ok := s.Get("a"); ok {
		t.Error("Get on empty store returned ok = true")
	}
	s.Put("a", "1")
	if v, ok := s.Get("a"); v != "1" || !ok {
		t.Errorf("Get = %q, %v, want \"1\", true", v, ok)
	}
	s.Put("a", "2")
	if v, _ := s.Get("a"); v != "2" {
		t.Errorf("after overwrite = %q, want \"2\"", v)
	}
}

func TestCopy(t *testing.T) {
	src, dst := NewMemStore(), NewMemStore()
	src.Put("a", "1")
	src.Put("b", "2")

	if got := Copy(src, dst, []string{"a", "zz"}); got != 1 {
		t.Errorf("Copy = %d, want 1", got)
	}
	if v, ok := dst.Get("a"); v != "1" || !ok {
		t.Errorf("dst[a] = %q, %v", v, ok)
	}
	if _, ok := dst.Get("zz"); ok {
		t.Error("missing key was copied")
	}
	if got := Copy(src, dst, nil); got != 0 {
		t.Errorf("Copy(nil keys) = %d, want 0", got)
	}
}
