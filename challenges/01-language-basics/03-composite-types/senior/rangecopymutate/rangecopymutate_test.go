package rangecopymutate

import "testing"

func TestDiscount(t *testing.T) {
	orders := []Order{{100}, {200}, {50}}
	Discount(orders, 10)
	want := []int{90, 180, 45}
	for i, w := range want {
		if orders[i].Price != w {
			t.Errorf("orders[%d].Price=%d; want %d", i, orders[i].Price, w)
		}
	}
}
