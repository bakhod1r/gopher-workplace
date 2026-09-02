package sliceleak

import "testing"

func TestPrefixAliases(t *testing.T) {
	src := make([]byte, 1<<20)
	src[0] = 'a'

	p := Prefix(src, 8)
	if len(p) != 8 {
		t.Errorf("len = %d, want 8", len(p))
	}
	if cap(p) != cap(src) {
		t.Errorf("cap = %d, want %d: Prefix must alias the source", cap(p), cap(src))
	}

	src[0] = 'z'
	if p[0] != 'z' {
		t.Error("Prefix should alias the source, not copy it")
	}
}

func TestPrefixCopyDetaches(t *testing.T) {
	src := make([]byte, 1<<20)
	src[0] = 'a'

	p := PrefixCopy(src, 8)
	if len(p) != 8 {
		t.Errorf("len = %d, want 8", len(p))
	}
	if cap(p) != 8 {
		t.Errorf("cap = %d, want 8: the copy must retain only what it keeps", cap(p))
	}

	src[0] = 'z'
	if p[0] != 'a' {
		t.Error("PrefixCopy must be independent of the source")
	}
}

func TestRetainedBytes(t *testing.T) {
	src := make([]byte, 1<<20)

	aliased := Prefix(src, 8)
	copied := PrefixCopy(src, 8)

	if RetainedBytes(aliased) < 1<<20 {
		t.Errorf("aliased retains %d bytes, want the whole array", RetainedBytes(aliased))
	}
	if RetainedBytes(copied) != 8 {
		t.Errorf("copied retains %d bytes, want 8", RetainedBytes(copied))
	}
	if RetainedBytes(copied) >= RetainedBytes(aliased) {
		t.Error("the copy must retain strictly less than the alias")
	}
}

func TestClamping(t *testing.T) {
	src := []byte("abc")

	if got := Prefix(src, 99); len(got) != 3 {
		t.Errorf("len = %d, want 3", len(got))
	}
	if got := PrefixCopy(src, 99); len(got) != 3 {
		t.Errorf("len = %d, want 3", len(got))
	}
	if got := Prefix(src, -1); len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
	if got := PrefixCopy(src, -1); len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}

func TestEmptySource(t *testing.T) {
	if got := Prefix(nil, 4); len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
	if got := PrefixCopy(nil, 4); len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}

func TestExtractorsDiffer(t *testing.T) {
	src := make([]byte, 1<<16)

	var a Extractor = Aliasing{}
	var c Extractor = Copying{}

	if RetainedBytes(a.Extract(src, 4)) <= RetainedBytes(c.Extract(src, 4)) {
		t.Error("the aliasing extractor must retain more than the copying one")
	}
}

func TestManyExtractsRetention(t *testing.T) {
	// 100 messages of 64KB, keeping 8 bytes of each.
	aliased := make([][]byte, 0, 100)
	copied := make([][]byte, 0, 100)
	for i := 0; i < 100; i++ {
		msg := make([]byte, 1<<16)
		aliased = append(aliased, Prefix(msg, 8))
		copied = append(copied, PrefixCopy(msg, 8))
	}

	aliasedTotal, copiedTotal := 0, 0
	for i := range aliased {
		aliasedTotal += RetainedBytes(aliased[i])
		copiedTotal += RetainedBytes(copied[i])
	}

	if copiedTotal >= aliasedTotal/100 {
		t.Errorf("copied retains %d bytes, aliased retains %d; the copy should be far smaller",
			copiedTotal, aliasedTotal)
	}
}

func BenchmarkPrefix(b *testing.B) {
	src := make([]byte, 1<<16)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Prefix(src, 8)
	}
}

func BenchmarkPrefixCopy(b *testing.B) {
	src := make([]byte, 1<<16)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = PrefixCopy(src, 8)
	}
}
