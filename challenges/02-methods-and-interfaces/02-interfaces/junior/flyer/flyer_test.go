package flyer

import "testing"

func TestFlyer(t *testing.T) {
	var f Flyer
	f = Bird{Species: "eagle"}
	if got := f.Fly(); got != "eagle flies" { t.Errorf("Bird = %q", got) }
	f = Plane{}
	if got := f.Fly(); got != "plane flies" { t.Errorf("Plane = %q", got) }
}
