package seqlockifc

import (
	"sync"
	"testing"
)

func TestReadBeforeWrite(t *testing.T) {
	var s SeqLock
	if got := s.Read(); got.Requests != 0 || got.Errors != 0 {
		t.Errorf("Read = %+v, want the zero snapshot", got)
	}
	if s.Seq() != 0 {
		t.Errorf("Seq = %d, want 0", s.Seq())
	}
}

func TestWriteThenRead(t *testing.T) {
	var s SeqLock
	s.Write(10, 2)
	if got := s.Read(); got.Requests != 10 || got.Errors != 2 {
		t.Errorf("Read = %+v, want {10 2}", got)
	}
}

func TestSequenceIsEvenWhenStable(t *testing.T) {
	var s SeqLock
	s.Write(1, 1)
	if s.Seq()%2 != 0 {
		t.Errorf("Seq = %d, want an even value when stable", s.Seq())
	}
	s.Write(2, 2)
	if s.Seq()%2 != 0 {
		t.Errorf("Seq = %d, want an even value when stable", s.Seq())
	}
}

func TestSequenceAdvancesPerWrite(t *testing.T) {
	var s SeqLock
	before := s.Seq()
	s.Write(1, 1)
	if s.Seq() <= before {
		t.Errorf("Seq did not advance: %d then %d", before, s.Seq())
	}
}

func TestNoTornReads(t *testing.T) {
	var s SeqLock
	stop := make(chan struct{})

	var writer sync.WaitGroup
	// The invariant: Errors is always exactly half of Requests.
	writer.Add(1)
	go func() {
		defer writer.Done()
		for i := int64(1); ; i++ {
			select {
			case <-stop:
				return
			default:
			}
			s.Write(i*2, i)
		}
	}()

	var readers sync.WaitGroup
	for r := 0; r < 4; r++ {
		readers.Add(1)
		go func() {
			defer readers.Done()
			for i := 0; i < 20000; i++ {
				snap := s.Read()
				if snap.Requests != snap.Errors*2 {
					t.Errorf("torn read: %+v", snap)
					return
				}
			}
		}()
	}

	readers.Wait()
	close(stop)
	writer.Wait()
}

func TestIsReader(t *testing.T) {
	var s SeqLock
	var r Reader = &s
	s.Write(3, 1)
	if got := r.Read(); got.Requests != 3 {
		t.Errorf("Read = %+v", got)
	}
}

func BenchmarkRead(b *testing.B) {
	var s SeqLock
	s.Write(1, 1)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = s.Read()
		}
	})
}
