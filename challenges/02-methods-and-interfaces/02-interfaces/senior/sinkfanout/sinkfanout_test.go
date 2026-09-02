package sinkfanout

import (
	"errors"
	"testing"
)

func TestMemSink(t *testing.T) {
	m := &MemSink{}
	if err := m.Write("a"); err != nil {
		t.Errorf("Write = %v, want nil", err)
	}
	if m.Len() != 1 {
		t.Errorf("Len = %d, want 1", m.Len())
	}
}

func TestErrSink(t *testing.T) {
	if err := (ErrSink{}).Write("a"); !errors.Is(err, ErrSinkFailed) {
		t.Errorf("Write = %v, want ErrSinkFailed", err)
	}
}

func TestFanOutCountsFailures(t *testing.T) {
	a, b := &MemSink{}, &MemSink{}
	got := FanOut([]Sink{a, ErrSink{}, b}, "e")
	if got != 1 {
		t.Errorf("FanOut = %d, want 1", got)
	}
	if a.Len() != 1 || b.Len() != 1 {
		t.Errorf("good sinks got %d and %d events, want 1 each", a.Len(), b.Len())
	}
}

func TestFanOutAllFail(t *testing.T) {
	if got := FanOut([]Sink{ErrSink{}, ErrSink{}}, "e"); got != 2 {
		t.Errorf("FanOut = %d, want 2", got)
	}
}

func TestFanOutEmpty(t *testing.T) {
	if got := FanOut(nil, "e"); got != 0 {
		t.Errorf("FanOut(nil) = %d, want 0", got)
	}
}

func TestFanOutManySinks(t *testing.T) {
	sinks := make([]Sink, 100)
	mems := make([]*MemSink, 100)
	for i := range sinks {
		mems[i] = &MemSink{}
		sinks[i] = mems[i]
	}

	if got := FanOut(sinks, "e"); got != 0 {
		t.Errorf("FanOut = %d, want 0", got)
	}
	for i, m := range mems {
		if m.Len() != 1 {
			t.Fatalf("sink %d got %d events, want 1", i, m.Len())
		}
	}
}

func TestFanOutRepeated(t *testing.T) {
	m := &MemSink{}
	for i := 0; i < 50; i++ {
		FanOut([]Sink{m, ErrSink{}}, "e")
	}
	if m.Len() != 50 {
		t.Errorf("Len = %d, want 50", m.Len())
	}
}

func BenchmarkFanOut(b *testing.B) {
	sinks := []Sink{&MemSink{}, &MemSink{}, ErrSink{}}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FanOut(sinks, "e")
	}
}
