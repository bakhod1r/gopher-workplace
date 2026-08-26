package toggle

import "testing"

func TestToggle(t *testing.T) {
	cases := []struct {
		name  string
		start bool
		want  bool
	}{
		{"off_to_on", false, true},
		{"on_to_off", true, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Switch{On: tc.start}
			s.Toggle()
			if s.On != tc.want {
				t.Errorf("Switch{%v}.Toggle() => On = %v, want %v",
					tc.start, s.On, tc.want)
			}
		})
	}

	// Double toggle should return to original.
	t.Run("double_toggle", func(t *testing.T) {
		s := Switch{On: false}
		s.Toggle()
		s.Toggle()
		if s.On != false {
			t.Errorf("double Toggle: On = %v, want false", s.On)
		}
	})

	// Triple toggle.
	t.Run("triple_toggle", func(t *testing.T) {
		s := Switch{On: true}
		s.Toggle()
		s.Toggle()
		s.Toggle()
		if s.On != false {
			t.Errorf("triple Toggle from true: On = %v, want false", s.On)
		}
	})
}
