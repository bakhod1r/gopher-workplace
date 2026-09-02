package zipmapgen

import "testing"

func TestZipMap(t *testing.T) {
	got := ZipMap([]string{"a", "b"}, []int{1, 2})
	if len(got) != 2 || got["a"] != 1 || got["b"] != 2 {
		t.Errorf("ZipMap = %v, want {a:1 b:2}", got)
	}
}

func TestZipMapStopsAtShorter(t *testing.T) {
	got := ZipMap([]string{"a", "b", "c"}, []int{1})
	if len(got) != 1 || got["a"] != 1 {
		t.Errorf("ZipMap = %v, want {a:1}", got)
	}
	got = ZipMap([]string{"a"}, []int{1, 2, 3})
	if len(got) != 1 || got["a"] != 1 {
		t.Errorf("ZipMap = %v, want {a:1}", got)
	}
}

func TestZipMapDuplicateKeys(t *testing.T) {
	got := ZipMap([]string{"a", "a"}, []int{1, 2})
	if len(got) != 1 || got["a"] != 2 {
		t.Errorf("ZipMap = %v, want {a:2} (last wins)", got)
	}
}

func TestZipMapEmpty(t *testing.T) {
	got := ZipMap([]string{"a"}, []int{})
	if got == nil {
		t.Fatal("ZipMap = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("ZipMap = %v, want {}", got)
	}
}
