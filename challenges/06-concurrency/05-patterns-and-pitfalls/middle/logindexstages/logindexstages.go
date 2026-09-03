// Package logindexstages — Gopher Workplace challenge.
package logindexstages

// Record is one parsed log line on its way to the search index.
type Record struct {
	Level   string
	Message string
}

// IndexLogs streams raw log lines through a three-stage pipeline — read,
// parse-and-filter, index — and returns the index IDs in input order.
//
// Every stage selects on done as well as on its channels, so closing done
// tears the whole pipeline down: no stage is left blocked on a send that
// nobody will ever receive.
//
// Examples:
//
//	IndexLogs(open done, ["err disk", "info ok"], parse, index)  => ["idx:disk"]
//	IndexLogs(open done, ["info ok"], parse, index)              => nil
//	IndexLogs(closed done, ["err disk"], parse, index)           => nil
func IndexLogs(done <-chan struct{}, lines []string, parse func(string) (Record, bool), index func(Record) string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
