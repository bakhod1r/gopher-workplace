// Package sitemapcrawler — Gopher Workplace challenge.
package sitemapcrawler

// FetchStatuses fetches every URL and reports the HTTP status codes in order.
//
// Examples:
//
//	FetchStatuses([]string{"/", "/gone"}, get)  => [200 404]
//	FetchStatuses([]string{"/"}, get)           => [200]
//	FetchStatuses(nil, get)                     => []
func FetchStatuses(urls []string, get func(url string) int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
