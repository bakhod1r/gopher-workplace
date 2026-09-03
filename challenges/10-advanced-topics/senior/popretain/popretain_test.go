package popretain

import "testing"

func TestPop(t *testing.T) {
	a, b := &Job{ID: 1}, &Job{ID: 2}
	got, rest := Pop([]*Job{a, b})
	if got != b {
		t.Errorf("Pop returned %v, want the last job", got)
	}
	if len(rest) != 1 || rest[0] != a {
		t.Errorf("rest = %v, want [a]", rest)
	}
}

func TestPopEmpty(t *testing.T) {
	got, rest := Pop(nil)
	if got != nil || len(rest) != 0 {
		t.Errorf("Pop(nil) = %v, %v, want nil, empty", got, rest)
	}
}

func TestPopClearsTheSlot(t *testing.T) {
	s := []*Job{{ID: 1}, {ID: 2}}
	_, rest := Pop(s)
	if s[1] != nil {
		t.Error("the popped slot still holds the job: it stays reachable through the backing array")
	}
	if len(rest) != 1 {
		t.Errorf("len = %d, want 1", len(rest))
	}
}

func TestPopRepeatedly(t *testing.T) {
	s := make([]*Job, 8)
	for i := range s {
		s[i] = &Job{ID: i}
	}
	backing := s
	for i := 7; i >= 0; i-- {
		var got *Job
		got, s = Pop(s)
		if got == nil || got.ID != i {
			t.Fatalf("pop %d returned %v", i, got)
		}
	}
	for i, p := range backing[:8] {
		if p != nil {
			t.Fatalf("slot %d still holds job %d after every pop", i, p.ID)
		}
	}
}
