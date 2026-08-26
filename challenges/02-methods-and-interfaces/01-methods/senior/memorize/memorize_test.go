package memorize

import "testing"

func TestMemoizer(t *testing.T) {
	calls := 0
	fn := func(k string) string {
		calls++
		return k + "-val"
	}

	m := New(fn)

	if got := m.Get("a"); got != "a-val" || calls != 1 {
		t.Errorf("first call failed: %q, %d", got, calls)
	}

	if got := m.Get("a"); got != "a-val" || calls != 1 {
		t.Errorf("second call should be cached: %q, %d", got, calls)
	}
}
