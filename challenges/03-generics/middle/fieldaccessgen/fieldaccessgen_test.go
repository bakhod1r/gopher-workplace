package fieldaccessgen

import "testing"

func TestTotalPrice(t *testing.T) {
	books := []Book{{200}, {350}}
	if got := TotalPrice(books, func(b Book) int { return b.Price }); got != 550 {
		t.Errorf("TotalPrice(books) = %d, want 550", got)
	}
	coffees := []Coffee{{450}}
	if got := TotalPrice(coffees, func(c Coffee) int { return c.Price }); got != 450 {
		t.Errorf("TotalPrice(coffees) = %d, want 450", got)
	}
	if got := TotalPrice([]Book{}, func(b Book) int { return b.Price }); got != 0 {
		t.Errorf("TotalPrice(empty) = %d, want 0", got)
	}
}

func TestTotalPriceArbitraryProjection(t *testing.T) {
	if got := TotalPrice([]string{"ab", "c"}, func(s string) int { return len(s) }); got != 3 {
		t.Errorf("TotalPrice = %d, want 3", got)
	}
}
