package pricefetcher

import (
	"reflect"
	"testing"
)

func TestFetchAllPrices(t *testing.T) {
	catalog := func(sku string) int { return len(sku) * 100 }

	cases := []struct {
		name        string
		skus        []string
		fetch       func(string) int
		discountPct int
		want        []int
	}{
		{"no_discount", []string{"ab", "cde"}, catalog, 0, []int{200, 300}},
		{"half_off", []string{"ab"}, catalog, 50, []int{100}},
		{"empty", []string{}, catalog, 10, []int{}},
		{"full_discount", []string{"ab", "cde"}, catalog, 100, []int{0, 0}},
		{"ten_percent", []string{"abcde"}, catalog, 10, []int{450}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FetchAllPrices(tc.skus, tc.fetch, tc.discountPct); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FetchAllPrices(%v) = %v, want %v", tc.skus, got, tc.want)
			}
		})
	}
}
