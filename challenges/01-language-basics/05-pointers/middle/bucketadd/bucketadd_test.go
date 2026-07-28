package bucketadd

import "testing"

func TestAddScore(t *testing.T) {
	m := map[string]*Bucket{}
	AddScore(m, "a", 3)
	AddScore(m, "a", 4)
	AddScore(m, "b", 5)
	if m["a"].Total != 7 || m["b"].Total != 5 {
		t.Errorf("a=%d b=%d want 7,5", m["a"].Total, m["b"].Total)
	}
}
