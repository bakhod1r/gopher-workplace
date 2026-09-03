package unionfindbug

import "testing"

func TestUnionDirect(t *testing.T) {
	var u UF[int]
	u.Union(1, 2)
	if !u.Connected(1, 2) {
		t.Error("Connected(1, 2) = false, want true")
	}
}

func TestUnionChain(t *testing.T) {
	var u UF[int]
	u.Union(1, 2)
	u.Union(1, 3)
	u.Union(4, 3)
	if !u.Connected(2, 3) {
		t.Error("Connected(2, 3) = false, want true")
	}
	if !u.Connected(1, 3) {
		t.Error("Connected(1, 3) = false, want true")
	}
	if !u.Connected(1, 4) {
		t.Error("Connected(1, 4) = false, want true")
	}
}

func TestUnionSeparate(t *testing.T) {
	var u UF[int]
	u.Union(1, 2)
	if u.Connected(1, 9) {
		t.Error("Connected(1, 9) = true, want false")
	}
}
