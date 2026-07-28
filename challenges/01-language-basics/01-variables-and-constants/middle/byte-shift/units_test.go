package units

import "testing"

func TestConsts(t *testing.T) {
	if KB != 1024 || MB != 1024*1024 || GB != 1024*1024*1024 {
		t.Fatalf("KB,MB,GB = %d,%d,%d", KB, MB, GB)
	}
	if TB != 1024*GB {
		t.Fatalf("TB = %d; want %d", TB, 1024*GB)
	}
}

func TestHumanize(t *testing.T) {
	cases := []struct {
		n     ByteSize
		count ByteSize
		sym   string
	}{
		{512, 512, "B"},
		{2 * KB, 2, "KB"},
		{3 * MB, 3, "MB"},
		{1 * GB, 1, "GB"},
		{5 * TB, 5, "TB"},
	}
	for _, c := range cases {
		gc, gs := Humanize(c.n)
		if gc != c.count || gs != c.sym {
			t.Errorf("Humanize(%d)=%d,%s; want %d,%s", c.n, gc, gs, c.count, c.sym)
		}
	}
}
