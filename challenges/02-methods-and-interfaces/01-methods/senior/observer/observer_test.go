package observer

import "testing"

func TestObserver(t *testing.T) {
	s := &Subject{}

	sum := 0
	s.Attach(func(state int) { sum += state })
	s.Attach(func(state int) { sum += state * 2 })

	s.SetState(10)
	if sum != 30 { // 10 + 20
		t.Errorf("sum = %d, want 30", sum)
	}
}
