package compareifc

import "testing"

func TestCompareTo(t *testing.T) {
	cases := []struct {
		name string
		a, b Score
		sign int
	}{
		{"greater", 5, 3, 1},
		{"equal", 2, 2, 0},
		{"less", 1, 4, -1},
		{"negatives", -5, -1, -1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a.CompareTo(tc.b)
			switch {
			case tc.sign == 0 && got != 0:
				t.Errorf("CompareTo = %d, want 0", got)
			case tc.sign > 0 && got <= 0:
				t.Errorf("CompareTo = %d, want >0", got)
			case tc.sign < 0 && got >= 0:
				t.Errorf("CompareTo = %d, want <0", got)
			}
		})
	}
}

func TestMax(t *testing.T) {
	if got := Max(Score(1), Score(9)).(Score); got != 9 {
		t.Errorf("Max = %d, want 9", got)
	}
	if got := Max(Score(9), Score(1)).(Score); got != 9 {
		t.Errorf("Max = %d, want 9", got)
	}
	if got := Max(Score(4), Score(4)).(Score); got != 4 {
		t.Errorf("Max = %d, want 4", got)
	}
}
