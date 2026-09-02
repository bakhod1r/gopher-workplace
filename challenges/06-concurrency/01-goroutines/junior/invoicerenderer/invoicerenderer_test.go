package invoicerenderer

import (
	"reflect"
	"testing"
)

func TestInvoiceTotals(t *testing.T) {
	cases := []struct {
		name     string
		invoices []Invoice
		want     []int
	}{
		{"one_invoice", []Invoice{{[]int{100, 250}}}, []int{350}},
		{"no_lines", []Invoice{{nil}}, []int{0}},
		{"empty", []Invoice{}, []int{}},
		{"credit_note", []Invoice{{[]int{500, -200}}}, []int{300}},
		{"batch", []Invoice{{[]int{1}}, {[]int{2, 3, 4}}, {[]int{}}}, []int{1, 9, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := InvoiceTotals(tc.invoices); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("InvoiceTotals(%v) = %v, want %v", tc.invoices, got, tc.want)
			}
		})
	}
}
