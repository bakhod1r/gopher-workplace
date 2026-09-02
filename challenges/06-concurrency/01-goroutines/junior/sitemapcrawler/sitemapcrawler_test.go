package sitemapcrawler

import (
	"reflect"
	"testing"
)

func TestFetchStatuses(t *testing.T) {
	get := func(url string) int {
		switch url {
		case "/":
			return 200
		case "/gone":
			return 404
		case "/boom":
			return 500
		}
		return 301
	}

	cases := []struct {
		name string
		urls []string
		get  func(string) int
		want []int
	}{
		{"ok_and_missing", []string{"/", "/gone"}, get, []int{200, 404}},
		{"single", []string{"/"}, get, []int{200}},
		{"empty", []string{}, get, []int{}},
		{"server_error", []string{"/boom", "/"}, get, []int{500, 200}},
		{"unknown_redirects", []string{"/a", "/b", "/gone"}, get, []int{301, 301, 404}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FetchStatuses(tc.urls, tc.get); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FetchStatuses(%v) = %v, want %v", tc.urls, got, tc.want)
			}
		})
	}
}
