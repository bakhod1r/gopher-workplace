package animalsay

import "testing"

func TestMakeNoise(t *testing.T) {
	if got := MakeNoise(Dog{}); got != "Woof!" {
		t.Errorf("Dog = %q, want \"Woof!\"", got)
	}
	if got := MakeNoise(Cat{}); got != "Meow!" {
		t.Errorf("Cat = %q, want \"Meow!\"", got)
	}
}

func TestChorus(t *testing.T) {
	cases := []struct {
		name string
		as   []Animal
		want string
	}{
		{"two", []Animal{Dog{}, Cat{}}, "Woof! Meow!"},
		{"one", []Animal{Cat{}}, "Meow!"},
		{"empty", nil, ""},
		{"three", []Animal{Dog{}, Dog{}, Cat{}}, "Woof! Woof! Meow!"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Chorus(tc.as); got != tc.want {
				t.Errorf("Chorus = %q, want %q", got, tc.want)
			}
		})
	}
}
