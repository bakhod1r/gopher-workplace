// Package csvimport — Gopher Workplace challenge.
package csvimport

// ImportCSV streams rows through a cleaning stage and a validation stage,
// returning the accepted rows in input order.
//
// Examples:
//
//	ImportCSV([]string{" a ", "  "}, trim, nonEmpty)  => []string{"a"}
//	ImportCSV([]string{" a ", " b "}, trim, nonEmpty)  => []string{"a", "b"}
//	ImportCSV(nil, trim, nonEmpty)                     => nil
func ImportCSV(rows []string, clean func(string) string, valid func(string) bool) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
