package idvalidate

import "testing"

func TestValidID(t *testing.T) {
	if !ValidID(UserID("u_1"), "u_") {
		t.Error(`ValidID(UserID("u_1"), "u_") = false, want true`)
	}
	if ValidID(UserID("x_1"), "u_") {
		t.Error(`ValidID(UserID("x_1"), "u_") = true, want false`)
	}
	if ValidID(UserID(""), "u_") {
		t.Error(`ValidID(UserID(""), "u_") = true, want false`)
	}
	if ValidID(UserID(""), "") {
		t.Error(`ValidID(UserID(""), "") = true, want false (empty is never valid)`)
	}
	if !ValidID(OrderID("o_9"), "o_") {
		t.Error(`ValidID(OrderID("o_9"), "o_") = false, want true`)
	}
	if !ValidID("plain", "pl") {
		t.Error(`ValidID("plain", "pl") = false, want true`)
	}
}
