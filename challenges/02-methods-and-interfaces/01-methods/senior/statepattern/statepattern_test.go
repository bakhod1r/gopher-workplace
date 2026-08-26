package statepattern

import "testing"

func TestState(t *testing.T) {
	d := &Document{State: Draft}
	d.Publish()
	if d.State != Moderation {
		t.Errorf("expected Moderation")
	}
	d.Publish()
	if d.State != Published {
		t.Errorf("expected Published")
	}
	d.Publish()
	if d.State != Published {
		t.Errorf("expected Published")
	}
}
