# Build It Once, However Many Ask

## Intuition

Lazy initialisation is a race unless the check and the build are one indivisible step. `sync.Once` provides exactly that, and it also provides the memory ordering that makes the built map safe to read afterwards.

## Approach

1. `t.once.Do(t.build)`.
2. Read `t.index[k]` and return the comma-ok result.

## Solution

```go
import (
	"sync"
	"sync/atomic"
)

// Builds counts how many times the index has been constructed.
var Builds atomic.Int64

// Table indexes a slice of pairs lazily.
type Table struct {
	once  sync.Once
	pairs [][2]string
	index map[string]int
}

// NewTable returns a table over pairs, without building the index.
func NewTable(pairs [][2]string) *Table {
	return &Table{pairs: pairs}
}

// build constructs the index. It must run at most once per table.
func (t *Table) build() {
	Builds.Add(1)
	t.index = make(map[string]int, len(t.pairs))
	for i, p := range t.pairs {
		t.index[p[0]] = i
	}
}

// Lookup returns the value for k, building the table's index on first use.
//
// The index is expensive, the callers are concurrent, and it must be built
// exactly once — every later lookup should be a plain map read.
//
// Examples:
//
// 	t := NewTable(pairs); t.Lookup("a") => the value for "a"
func (t *Table) Lookup(k string) (int, bool) {
	t.once.Do(t.build)
	i, ok := t.index[k]
	return i, ok
}
```

## Walkthrough

Under 32 concurrent first calls, one goroutine runs `build` and the other 31 block inside `Do` until it returns — so every one of them reads a complete index.

## Pitfalls

- `if t.index == nil { t.build() }` — the race the `Once` exists to remove.
- Copying the `Table` after first use; a `sync.Once` must not be copied.
- Putting the `Once` in a package-level variable, which would build only the first table's index.
