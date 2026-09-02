package fmtstring

import (
	"fmt"
	"testing"
)

func TestMoneyString(t *testing.T) {
	cases := []struct {
		m    Money
		want string
	}{
		{1234, "12.34"},
		{5, "0.05"},
		{0, "0.00"},
		{100, "1.00"},
		{-250, "-2.50"},
	}

	for _, tc := range cases {
		if got := tc.m.String(); got != tc.want {
			t.Errorf("Money(%d).String() = %q, want %q", int(tc.m), got, tc.want)
		}
	}
}

func TestLevelString(t *testing.T) {
	if got := Debug.String(); got != "DEBUG" {
		t.Errorf("Debug = %q", got)
	}
	if got := Error.String(); got != "ERROR" {
		t.Errorf("Error = %q", got)
	}
	if got := Level(7).String(); got != "LEVEL(7)" {
		t.Errorf("Level(7) = %q, want \"LEVEL(7)\"", got)
	}
}

func TestLine(t *testing.T) {
	if got := Line(Info, "paid", Money(1234)); got != "[INFO] paid: 12.34" {
		t.Errorf("Line = %q", got)
	}
	if got := Line(Level(9), "odd", Money(-5)); got != "[LEVEL(9)] odd: -0.05" {
		t.Errorf("Line = %q", got)
	}
}

func TestUsedByFmt(t *testing.T) {
	if got := fmt.Sprintf("%v %s", Money(1), Error); got != "0.01 ERROR" {
		t.Errorf("fmt = %q, want \"0.01 ERROR\"", got)
	}
}
