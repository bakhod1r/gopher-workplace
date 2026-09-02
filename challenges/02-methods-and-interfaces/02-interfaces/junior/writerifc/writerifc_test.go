package writerifc

import "testing"

func TestWrite(t *testing.T) {
	b := &Builder{}
	if n := b.Write("ab"); n != 2 {
		t.Errorf("Write = %d, want 2", n)
	}
	if n := b.Write(""); n != 0 {
		t.Errorf("Write(\"\") = %d, want 0", n)
	}
	if got := b.String(); got != "ab" {
		t.Errorf("String = %q, want \"ab\"", got)
	}
}

func TestWriteLines(t *testing.T) {
	b := &Builder{}
	if n := WriteLines(b, []string{"a", "b"}); n != 4 {
		t.Errorf("WriteLines = %d, want 4", n)
	}
	if got := b.String(); got != "a\nb\n" {
		t.Errorf("String = %q, want \"a\\nb\\n\"", got)
	}

	empty := &Builder{}
	if n := WriteLines(empty, nil); n != 0 {
		t.Errorf("WriteLines(nil) = %d, want 0", n)
	}
	if got := empty.String(); got != "" {
		t.Errorf("String = %q, want empty", got)
	}
}
