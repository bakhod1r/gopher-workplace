// Package pricefetcher — Gopher Workplace challenge.
package pricefetcher

// FetchAllPrices returns the discounted price of every SKU, in input order.
//
// Examples:
//
//	FetchAllPrices([]string{"ab", "cde"}, catalog, 0)   => [200 300]
//	FetchAllPrices([]string{"ab"}, catalog, 50)         => [100]
//	FetchAllPrices(nil, catalog, 10)                    => []
func FetchAllPrices(skus []string, fetch func(sku string) int, discountPct int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
