// Package logcounts — Gopher Workplace challenge.
package logcounts

// CountLevels counts log lines by severity level, processing each chunk in
// its own goroutine and merging the per-chunk counts under a mutex.
//
// Examples:
//
//	CountLevels([][]string{{"ERR a"}, {"ERR b"}}, first)  => map[ERR:2]
//	CountLevels([][]string{{"ERR a", "INFO b"}}, first)   => map[ERR:1 INFO:1]
//	CountLevels(nil, first)                               => empty map
func CountLevels(chunks [][]string, level func(string) string) map[string]int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
