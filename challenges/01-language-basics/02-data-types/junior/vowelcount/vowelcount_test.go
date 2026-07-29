package vowelcount

import "testing"

func TestVowels(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"hello", 2},
		{"AEIOU", 5},
		{"xyz", 0},
		{"café", 1}, // é is not an ASCII vowel
		{"", 0},
	}
	for _, c := range cases {
		if got := Vowels(c.s); got != c.want {
			t.Errorf("Vowels(%q)=%d; want %d", c.s, got, c.want)
		}
	}
}
