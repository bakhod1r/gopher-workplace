package rawpath

import "testing"

func TestTempPath(t *testing.T) {
	want := "C:\\Users\\temp\\log.txt"
	if got := TempPath(); got != want {
		t.Errorf("TempPath()=%q; want %q", got, want)
	}
}
