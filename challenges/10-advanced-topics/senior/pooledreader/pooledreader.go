// Package pooledreader — Gopher Workplace challenge.
package pooledreader

import (
	"bufio"
	"io"
	"strings"
	"sync"
)

var pool = sync.Pool{New: func() any { return bufio.NewReaderSize(nil, 64) }}

// FirstLine returns the first '\n'-terminated line from r, using a
// pooled bufio.Reader.
//
// A bufio.Reader taken from a pool is still attached to the previous
// source, with the previous source's buffered bytes inside it.
//
// Examples:
//
//	FirstLine(strings.NewReader("a\nb\n")) => "a", nil
func FirstLine(r io.Reader) (string, error) {
	// CHANGE CODE BELOW THIS LINE
	br := pool.Get().(*bufio.Reader)
	defer pool.Put(br)
	line, err := br.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSuffix(line, "\n"), nil
	// CHANGE CODE ABOVE THIS LINE
}
