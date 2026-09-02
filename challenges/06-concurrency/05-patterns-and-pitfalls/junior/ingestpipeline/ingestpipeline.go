// Package ingestpipeline — Gopher Workplace challenge.
package ingestpipeline

// IngestPipeline streams raw log lines through a parse stage and an error
// filter stage, returning the error records in input order.
//
// Examples:
//
//	IngestPipeline([]string{"err a", "info b"}, upper, hasErrPrefix)  => []string{"ERR A"}
//	IngestPipeline(only info lines, ...)                             => nil
//	IngestPipeline(nil, ...)                                         => nil
func IngestPipeline(lines []string, parse func(string) string, isError func(string) bool) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
