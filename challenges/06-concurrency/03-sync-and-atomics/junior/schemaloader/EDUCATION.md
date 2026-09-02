# One-Shot Schema Load

## Intuition

`sync.Once` gives a happens-before edge: everything the winning goroutine wrote inside `Do` is visible to every caller that returns from `Do` later. That is what makes reading the cached fields afterwards safe without any extra lock.

## Approach

1. Hold `once sync.Once`, `schema string`, `err error`, `attempts int`.
2. In `Load`, do all the work and the bookkeeping inside `once.Do`.
3. Return `l.schema, l.err`.

## Solution

```go
// Package schemaloader - Gopher Workplace challenge.
package schemaloader

import "sync"

// Loader parses a GraphQL schema exactly once and caches the outcome.
type Loader struct {
	once     sync.Once
	parse    func() (string, error)
	schema   string
	err      error
	attempts int
}

// NewLoader returns a Loader that calls parse on first use.
func NewLoader(parse func() (string, error)) *Loader {
	return &Loader{parse: parse}
}

// Load returns the parsed schema, parsing on the first call only.
//
// Examples:
//
//	l := NewLoader(func() (string, error) { return "schema", nil }); l.Load() => "schema", nil
//	l.Load(); l.Load()                                                        => same pair, parsed once
func (l *Loader) Load() (string, error) {
	l.once.Do(func() {
		l.attempts++
		l.schema, l.err = l.parse()
	})
	return l.schema, l.err
}

// Attempts reports how many times the parse function ran.
//
// Examples:
//
//	l.Load(); l.Load(); l.Attempts() => 1
func (l *Loader) Attempts() int {
	return l.attempts
}
```

## Walkthrough

Twelve queries arrive at once. One goroutine wins the `Do`, parses, and stores the schema and a nil error. The other eleven block, then all return that same pair.

## Pitfalls

- Retrying inside `Load` when `err != nil` — `Once` will not run again, so the retry silently does nothing.
- Assigning the results outside the `Do` closure, which races.
- Returning early before `Do` completes, handing back a zero value.
