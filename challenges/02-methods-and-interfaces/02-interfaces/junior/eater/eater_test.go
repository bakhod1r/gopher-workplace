package eater

import "testing"

func TestEats(t *testing.T) {
	if !(Cow{}).Eats("grass") {
		t.Error("Cow should eat grass")
	}
	if (Cow{}).Eats("meat") {
		t.Error("Cow should not eat meat")
	}
	if !(Lion{}).Eats("meat") {
		t.Error("Lion should eat meat")
	}
	if (Lion{}).Eats("grass") {
		t.Error("Lion should not eat grass")
	}
}

func TestFeedableCount(t *testing.T) {
	cases := []struct {
		name string
		es   []Eater
		food string
		want int
	}{
		{"two_cows", []Eater{Cow{}, Lion{}, Cow{}}, "grass", 2},
		{"one_lion", []Eater{Cow{}, Lion{}}, "meat", 1},
		{"none", []Eater{Cow{}, Lion{}}, "fish", 0},
		{"empty", nil, "grass", 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FeedableCount(tc.es, tc.food); got != tc.want {
				t.Errorf("FeedableCount = %d, want %d", got, tc.want)
			}
		})
	}
}
