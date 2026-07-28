package weekday

import "testing"

func TestValues(t *testing.T) {
	want := []Day{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}
	for i, d := range want {
		if int(d) != i {
			t.Errorf("day %d = %d; want %d", i, d, i)
		}
	}
}

func TestWeekend(t *testing.T) {
	for d := Sunday; d <= Saturday; d++ {
		got := IsWeekend(d)
		want := d == Saturday || d == Sunday
		if got != want {
			t.Errorf("IsWeekend(%d)=%v; want %v", d, got, want)
		}
	}
}
