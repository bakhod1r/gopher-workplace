package logger

import "testing"

func TestMemLogger(t *testing.T) {
	m := &MemLogger{}
	m.Log("a")
	m.Log("b")
	if len(m.Lines) != 2 || m.Lines[0] != "a" || m.Lines[1] != "b" {
		t.Errorf("Lines = %v, want [a b]", m.Lines)
	}
}

func TestLogAll(t *testing.T) {
	t.Run("memory", func(t *testing.T) {
		m := &MemLogger{}
		LogAll(m, []string{"x", "y", "z"})
		if len(m.Lines) != 3 {
			t.Fatalf("len = %d, want 3", len(m.Lines))
		}
		if m.Lines[2] != "z" {
			t.Errorf("last = %q, want \"z\"", m.Lines[2])
		}
	})

	t.Run("discard", func(t *testing.T) {
		LogAll(Discard{}, []string{"a"})
	})

	t.Run("empty", func(t *testing.T) {
		m := &MemLogger{}
		LogAll(m, nil)
		if len(m.Lines) != 0 {
			t.Errorf("len = %d, want 0", len(m.Lines))
		}
	})
}
