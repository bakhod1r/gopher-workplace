package pipelinestg

import (
	"testing"
	"time"
)

func feed(vs ...int) <-chan int {
	ch := make(chan int, len(vs))
	for _, v := range vs {
		ch <- v
	}
	close(ch)
	return ch
}

func drain(ch <-chan int) []int {
	var out []int
	for v := range ch {
		out = append(out, v)
	}
	return out
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

func TestDoubleStage(t *testing.T) {
	if got := drain(RunStage(feed(1, 2), DoubleStage{})); !eq(got, []int{2, 4}) {
		t.Errorf("got %v, want [2 4]", got)
	}
}

func TestDropOddStage(t *testing.T) {
	if got := drain(RunStage(feed(1, 2, 3, 4), DropOddStage{})); !eq(got, []int{2, 4}) {
		t.Errorf("got %v, want [2 4]", got)
	}
}

func TestEmptyInputClosesOutput(t *testing.T) {
	out := RunStage(feed(), DoubleStage{})
	select {
	case _, ok := <-out:
		if ok {
			t.Error("expected a closed channel")
		}
	case <-time.After(time.Second):
		t.Fatal("output channel was never closed")
	}
}

func TestChainedStages(t *testing.T) {
	stage1 := RunStage(feed(1, 2, 3, 4), DoubleStage{})
	stage2 := RunStage(stage1, DropOddStage{})
	if got := drain(stage2); !eq(got, []int{2, 4, 6, 8}) {
		t.Errorf("got %v, want [2 4 6 8]", got)
	}
}

func TestAllDropped(t *testing.T) {
	if got := drain(RunStage(feed(1, 3, 5), DropOddStage{})); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestLargeStream(t *testing.T) {
	in := make(chan int)
	go func() {
		defer close(in)
		for i := 0; i < 10000; i++ {
			in <- i
		}
	}()

	n := 0
	for range RunStage(in, DoubleStage{}) {
		n++
	}
	if n != 10000 {
		t.Errorf("received %d values, want 10000", n)
	}
}
