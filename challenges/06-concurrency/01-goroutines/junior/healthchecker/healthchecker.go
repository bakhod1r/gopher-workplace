// Package healthchecker — Gopher Workplace challenge.
package healthchecker

// CheckAll reports which services answered with a healthy status code.
//
// Examples:
//
//	CheckAll([]string{"api", "db"}, probe)  => [true false]
//	CheckAll([]string{"api"}, probe)        => [true]
//	CheckAll(nil, probe)                    => []
func CheckAll(services []string, probe func(service string) int) []bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
