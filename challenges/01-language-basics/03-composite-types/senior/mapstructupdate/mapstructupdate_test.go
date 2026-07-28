package mapstructupdate

import "testing"

func TestRecord(t *testing.T) {
	m := map[string]Stat{}
	Record(m, "a")
	Record(m, "a")
	Record(m, "b")
	if m["a"].Hits != 2 || m["b"].Hits != 1 {
		t.Errorf("got a=%d b=%d; want 2,1", m["a"].Hits, m["b"].Hits)
	}
}
