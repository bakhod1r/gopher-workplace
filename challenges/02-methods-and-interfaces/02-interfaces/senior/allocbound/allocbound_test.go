package allocbound

import "testing"

func TestWriteAndLen(t *testing.T) {
	s := NewSink(4)
	s.Write(1)
	s.Write(2)
	s.Write(3)
	if s.Len() != 3 {
		t.Errorf("Len = %d, want 3", s.Len())
	}
}

func TestFill(t *testing.T) {
	s := NewSink(1000)
	Fill(s, 1000)
	if s.Len() != 1000 {
		t.Errorf("Len = %d, want 1000", s.Len())
	}
}

func TestFillZero(t *testing.T) {
	s := NewSink(0)
	Fill(s, 0)
	if s.Len() != 0 {
		t.Errorf("Len = %d, want 0", s.Len())
	}
}

func TestNoAllocsWithinCapacity(t *testing.T) {
	s := NewSink(4096)
	var r Recorder = s

	avg := testing.AllocsPerRun(100, func() {
		if s.Len() >= 4096 {
			s.buf = s.buf[:0]
		}
		r.Write(1)
	})
	if avg > 0 {
		t.Errorf("Write allocated %.2f times per call, want 0 within capacity", avg)
	}
}

func TestCapacityReserved(t *testing.T) {
	s := NewSink(512)
	if c := cap(s.buf); c < 512 {
		t.Errorf("cap = %d, want at least 512 reserved up front", c)
	}
	before := cap(s.buf)
	Fill(s, 512)
	if cap(s.buf) != before {
		t.Errorf("cap grew from %d to %d during Fill", before, cap(s.buf))
	}
}

func BenchmarkFill(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := NewSink(1024)
		Fill(s, 1024)
	}
}
