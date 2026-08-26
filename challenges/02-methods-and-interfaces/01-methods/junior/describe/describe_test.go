package describe

import "testing"

func TestCircleDescribe(t *testing.T) {
	cases := []struct {
		name string
		c    Circle
		want string
	}{
		{"integer", Circle{5}, "Circle(radius=5)"},
		{"decimal", Circle{3.5}, "Circle(radius=3.5)"},
		{"zero", Circle{0}, "Circle(radius=0)"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.c.Describe(); got != tc.want {
				t.Errorf("Circle{%g}.Describe() = %q, want %q",
					tc.c.Radius, got, tc.want)
			}
		})
	}
}

func TestSquareDescribe(t *testing.T) {
	cases := []struct {
		name string
		s    Square
		want string
	}{
		{"integer", Square{4}, "Square(side=4)"},
		{"decimal", Square{2.5}, "Square(side=2.5)"},
		{"zero", Square{0}, "Square(side=0)"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.s.Describe(); got != tc.want {
				t.Errorf("Square{%g}.Describe() = %q, want %q",
					tc.s.Side, got, tc.want)
			}
		})
	}
}
