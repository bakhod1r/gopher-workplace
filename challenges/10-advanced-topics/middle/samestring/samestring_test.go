package samestring

import (
	"strings"
	"testing"
)

func TestSameBytesIdentical(t *testing.T) {
	s := strings.Repeat("x", 32)
	if !SameBytes(s, s) {
		t.Error("SameBytes(s, s) = false, want true")
	}
	if !SameBytes(s, s[:32]) {
		t.Error("SameBytes(s, s[:32]) = false, want true")
	}
}

func TestSameBytesSubstrings(t *testing.T) {
	s := strings.Repeat("y", 8)
	if SameBytes(s, s[:4]) {
		t.Error("different lengths must report false")
	}
	if SameBytes(s[1:], s[:7]) {
		t.Error("different starts must report false")
	}
}

func TestSameBytesSeparateCopies(t *testing.T) {
	a := strings.Repeat("z", 16)
	b := string([]byte(a))
	if a != b {
		t.Fatal("the fixture is wrong: the strings must be equal")
	}
	if SameBytes(a, b) {
		t.Error("two separate copies must report false")
	}
}

func TestSameBytesEmpty(t *testing.T) {
	if SameBytes("", "") {
		t.Error("empty strings have no storage to share")
	}
	if SameBytes("a", "") {
		t.Error("different lengths must report false")
	}
}
