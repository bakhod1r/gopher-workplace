package notifier

import "testing"

func TestEmailNotify(t *testing.T) {
	e := &Email{}
	if !e.Notify("hi") {
		t.Error("Email.Notify = false, want true")
	}
	if len(e.Sent) != 1 || e.Sent[0] != "hi" {
		t.Errorf("Sent = %v, want [hi]", e.Sent)
	}
}

func TestBroadcast(t *testing.T) {
	a, b := &Email{}, &Email{}
	if got := Broadcast([]Notifier{a, Broken{}, b}, "alert"); got != 2 {
		t.Errorf("Broadcast = %d, want 2", got)
	}
	if len(b.Sent) != 1 {
		t.Error("Broadcast stopped early: last notifier never called")
	}
	if got := Broadcast(nil, "x"); got != 0 {
		t.Errorf("Broadcast(nil) = %d, want 0", got)
	}
	if got := Broadcast([]Notifier{Broken{}}, "x"); got != 0 {
		t.Errorf("Broadcast = %d, want 0", got)
	}
}
