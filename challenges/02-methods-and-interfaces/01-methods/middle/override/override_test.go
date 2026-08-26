package override

import "testing"

func TestAnimalString(t *testing.T) {
	a := Animal{Species: "Dog"}
	want := "Animal(Dog)"
	if got := a.String(); got != want {
		t.Errorf("Animal.String() = %q, want %q", got, want)
	}
}

func TestPetString(t *testing.T) {
	cases := []struct {
		name string
		p    Pet
		want string
	}{
		{"cat", Pet{Animal{"Cat"}, "Whiskers"}, "Pet(Whiskers, Cat)"},
		{"dog", Pet{Animal{"Dog"}, "Rex"}, "Pet(Rex, Dog)"},
		{"empty", Pet{Animal{""}, ""}, "Pet(, )"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.p.String(); got != tc.want {
				t.Errorf("Pet.String() = %q, want %q", got, tc.want)
			}
		})
	}
}
