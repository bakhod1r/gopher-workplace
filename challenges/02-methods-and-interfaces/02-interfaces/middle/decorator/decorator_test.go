package decorator

import "testing"

func TestBase(t *testing.T) {
	if got := (Base{Amount: 100}).Price(); got != 100 {
		t.Errorf("Base.Price = %d, want 100", got)
	}
}

func TestDiscount(t *testing.T) {
	if got := (Discount{Inner: Base{Amount: 100}, Percent: 10}).Price(); got != 90 {
		t.Errorf("Price = %d, want 90", got)
	}
	if got := (Discount{Inner: Base{Amount: 99}, Percent: 33}).Price(); got != 67 {
		t.Errorf("Price = %d, want 67", got)
	}
	if got := (Discount{Inner: Base{Amount: 100}, Percent: 0}).Price(); got != 100 {
		t.Errorf("Price = %d, want 100", got)
	}
}

func TestTaxOnDiscount(t *testing.T) {
	p := Tax{Inner: Discount{Inner: Base{Amount: 100}, Percent: 10}, Percent: 20}
	if got := p.Price(); got != 108 {
		t.Errorf("Price = %d, want 108", got)
	}
}

func TestWrap(t *testing.T) {
	p := Wrap(Base{Amount: 100},
		func(in Pricer) Pricer { return Discount{Inner: in, Percent: 10} },
		func(in Pricer) Pricer { return Tax{Inner: in, Percent: 20} },
	)
	if got := p.Price(); got != 108 {
		t.Errorf("Price = %d, want 108", got)
	}

	if got := Wrap(Base{Amount: 42}).Price(); got != 42 {
		t.Errorf("Wrap with no layers = %d, want 42", got)
	}
}
