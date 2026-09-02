package stringer

import (
	"fmt"
	"testing"
)

func TestColorString(t *testing.T) {
	cases := []struct {
		c    Color
		want string
	}{
		{Red, "red"},
		{Green, "green"},
		{Blue, "blue"},
		{Color(9), "unknown"},
		{Color(-1), "unknown"},
	}

	for _, tc := range cases {
		if got := tc.c.String(); got != tc.want {
			t.Errorf("Color(%d).String() = %q, want %q", int(tc.c), got, tc.want)
		}
	}
}

func TestTempString(t *testing.T) {
	if got := Temp(21).String(); got != "21C" {
		t.Errorf("Temp.String = %q, want \"21C\"", got)
	}
	if got := Temp(-5).String(); got != "-5C" {
		t.Errorf("Temp.String = %q, want \"-5C\"", got)
	}
}

func TestPrint(t *testing.T) {
	if got := Print(Temp(21)); got != "21C" {
		t.Errorf("Print = %q", got)
	}
	if got := fmt.Sprintf("%v", Green); got != "green" {
		t.Errorf("fmt %%v = %q, want \"green\"", got)
	}
}
