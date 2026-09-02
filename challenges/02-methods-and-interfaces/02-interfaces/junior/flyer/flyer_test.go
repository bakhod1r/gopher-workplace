package flyer

import "testing"

func TestAltitude(t *testing.T) {
	if got := (Bird{Meters: 100}).Altitude(); got != 100 {
		t.Errorf("Bird.Altitude = %d, want 100", got)
	}
	if got := (Jet{Feet: 30000}).Altitude(); got != 9144 {
		t.Errorf("Jet.Altitude = %d, want 9144", got)
	}
	if got := (Jet{Feet: 1000}).Altitude(); got != 304 {
		t.Errorf("Jet.Altitude = %d, want 304", got)
	}
}

func TestHighest(t *testing.T) {
	cases := []struct {
		name string
		fs   []Flyer
		want int
	}{
		{"jet_wins", []Flyer{Bird{Meters: 100}, Jet{Feet: 1000}}, 304},
		{"bird_wins", []Flyer{Bird{Meters: 500}, Jet{Feet: 1000}}, 500},
		{"empty", nil, 0},
		{"single", []Flyer{Jet{Feet: 30000}}, 9144},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Highest(tc.fs); got != tc.want {
				t.Errorf("Highest = %d, want %d", got, tc.want)
			}
		})
	}
}
