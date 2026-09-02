package describer

import "testing"

func TestDescribe(t *testing.T) {
	if got := (User{Name: "ann"}).Describe(); got != "user ann" {
		t.Errorf("User.Describe = %q", got)
	}
	if got := (Server{Host: "db", Port: 5432}).Describe(); got != "server db:5432" {
		t.Errorf("Server.Describe = %q", got)
	}
}

func TestDescribeAll(t *testing.T) {
	got := DescribeAll([]Describer{User{Name: "a"}, Server{Host: "h", Port: 80}})
	want := []string{"user a", "server h:80"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if n := len(DescribeAll(nil)); n != 0 {
		t.Errorf("DescribeAll(nil) len = %d, want 0", n)
	}
}
