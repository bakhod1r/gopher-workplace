package statemach

import "testing"

func TestTransitions(t *testing.T) {
	next, ok := Pending{}.Next("ship")
	if !ok || next.Name() != "shipped" {
		t.Errorf("Pending ship => %s, %v", next.Name(), ok)
	}

	next, ok = Pending{}.Next("cancel")
	if ok || next.Name() != "pending" {
		t.Errorf("Pending cancel => %s, %v; want pending, false", next.Name(), ok)
	}

	next, ok = Shipped{}.Next("deliver")
	if !ok || next.Name() != "delivered" {
		t.Errorf("Shipped deliver => %s, %v", next.Name(), ok)
	}

	next, ok = Delivered{}.Next("ship")
	if ok || next.Name() != "delivered" {
		t.Errorf("Delivered ship => %s, %v; want delivered, false", next.Name(), ok)
	}
}

func TestRun(t *testing.T) {
	cases := []struct {
		name     string
		events   []string
		wantName string
		wantOK   int
	}{
		{"happy_path", []string{"ship", "deliver"}, "delivered", 2},
		{"bad_event", []string{"deliver", "ship"}, "shipped", 1},
		{"none", nil, "pending", 0},
		{"past_terminal", []string{"ship", "deliver", "ship"}, "delivered", 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotName, gotOK := Run(Pending{}, tc.events)
			if gotName != tc.wantName || gotOK != tc.wantOK {
				t.Errorf("Run = %q, %d; want %q, %d", gotName, gotOK, tc.wantName, tc.wantOK)
			}
		})
	}
}
