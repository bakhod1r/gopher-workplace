package joinv

import "testing"

func TestJoin(t *testing.T) {
	if got := Join(","); got != "" {
		t.Errorf("Join(\",\")=%q want empty", got)
	}
	if got := Join("-", "a"); got != "a" {
		t.Errorf("single=%q want a", got)
	}
	if got := Join("/", "a", "b", "c"); got != "a/b/c" {
		t.Errorf("=%q want a/b/c", got)
	}
}
