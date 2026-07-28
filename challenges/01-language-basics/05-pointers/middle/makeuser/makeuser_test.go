package makeuser

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser("ann", 30)
	if u == nil || u.Name != "ann" || u.Age != 30 {
		t.Errorf("bad user: %+v", u)
	}
	v := NewUser("ann", 30)
	if u == v {
		t.Errorf("each call should allocate a new instance")
	}
}
