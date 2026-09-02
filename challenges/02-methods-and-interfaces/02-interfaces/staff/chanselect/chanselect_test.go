package chanselect

import "testing"

func filled(vs ...int) <-chan int {
	ch := make(chan int, len(vs))
	for _, v := range vs {
		ch <- v
	}
	close(ch)
	return ch
}

func eq(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestTryRecvEmpty(t *testing.T) {
	ch := make(chan int)
	if v, ok, ready := TryRecv(ch); ready || ok || v != 0 {
		t.Errorf("TryRecv = %d, %v, %v; want 0, false, false", v, ok, ready)
	}
}

func TestTryRecvValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 7
	if v, ok, ready := TryRecv(ch); v != 7 || !ok || !ready {
		t.Errorf("TryRecv = %d, %v, %v; want 7, true, true", v, ok, ready)
	}
}

func TestTryRecvClosed(t *testing.T) {
	ch := make(chan int)
	close(ch)
	if v, ok, ready := TryRecv(ch); v != 0 || ok || !ready {
		t.Errorf("TryRecv = %d, %v, %v; want 0, false, true", v, ok, ready)
	}
}

func TestDrainBoth(t *testing.T) {
	got := Drain(filled(1, 3), filled(2, 4))
	if !eq(got, []int{1, 2, 3, 4}) {
		t.Errorf("Drain = %v, want [1 2 3 4]", got)
	}
}

func TestDrainOneEmpty(t *testing.T) {
	got := Drain(filled(), filled(5))
	if !eq(got, []int{5}) {
		t.Errorf("Drain = %v, want [5]", got)
	}
}

func TestDrainBothEmpty(t *testing.T) {
	if got := Drain(filled(), filled()); len(got) != 0 {
		t.Errorf("Drain = %v, want empty", got)
	}
}

func TestDrainLargeStreams(t *testing.T) {
	a := make(chan int, 500)
	b := make(chan int, 500)
	for i := 0; i < 500; i++ {
		a <- i * 2
		b <- i*2 + 1
	}
	close(a)
	close(b)

	got := Drain(a, b)
	if len(got) != 1000 {
		t.Fatalf("len = %d, want 1000", len(got))
	}
	for i := range got {
		if got[i] != i {
			t.Fatalf("got[%d] = %d, want %d", i, got[i], i)
		}
	}
}

func TestFirstReadyPicksOpenChannel(t *testing.T) {
	a := make(chan int) // never ready
	close(a)
	b := filled(9)

	v, from := FirstReady(a, b)
	if v != 9 || from != "b" {
		t.Errorf("FirstReady = %d, %q; want 9, \"b\"", v, from)
	}
}

func TestFirstReadyBothClosed(t *testing.T) {
	v, from := FirstReady(filled(), filled())
	if v != 0 || from != "" {
		t.Errorf("FirstReady = %d, %q; want 0, \"\"", v, from)
	}
}

func TestIsSource(t *testing.T) {
	var s Source = ChanSource{C: filled(1)}
	if v, ok, ready := TryRecv(s.Chan()); v != 1 || !ok || !ready {
		t.Errorf("TryRecv = %d, %v, %v", v, ok, ready)
	}
}
