package animalsay

import "testing"

func TestAnimalSay(t *testing.T) {
	if got := MakeNoise(Dog{}); got != "Woof!" { t.Errorf("Dog = %q", got) }
	if got := MakeNoise(Cat{}); got != "Meow!" { t.Errorf("Cat = %q", got) }
}
