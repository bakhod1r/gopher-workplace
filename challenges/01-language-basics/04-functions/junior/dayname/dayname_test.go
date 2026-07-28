package dayname

import "testing"

func TestDayName(t *testing.T) {
	want := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	for i, w := range want {
		if got := DayName(i); got != w {
			t.Errorf("DayName(%d)=%q want %q", i, got, w)
		}
	}
	if DayName(9) != "?" {
		t.Errorf("out of range should be ?")
	}
}
