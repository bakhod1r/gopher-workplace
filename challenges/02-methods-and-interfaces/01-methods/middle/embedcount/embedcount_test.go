package embedcount

import "testing"

func TestJobRun(t *testing.T) {
	j := Job{Name: "task"}
	if j.Count != 0 {
		t.Fatalf("initial count = %d, want 0", j.Count)
	}

	j.Run()
	if j.Count != 1 {
		t.Errorf("after Run(): count = %d, want 1", j.Count)
	}

	j.Run()
	if j.Count != 2 {
		t.Errorf("after second Run(): count = %d, want 2", j.Count)
	}
}
