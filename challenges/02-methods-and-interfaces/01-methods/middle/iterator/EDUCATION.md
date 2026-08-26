# Iterator Pattern

## Intuition

The `Next()/Value()` pattern is Go's standard iterator idiom (see
`bufio.Scanner`, `sql.Rows`). `Next` moves forward and returns whether
there's more data. `Value` reads the current position.

## Approach

1. `Next`: if `pos < len(data)`, increment `pos`, return true.
2. `Value`: return `data[pos-1]`.

## Solution

```go
func (it *IntIter) Next() bool {
	if it.pos >= len(it.data) {
		return false
	}
	it.pos++
	return true
}

func (it *IntIter) Value() int {
	return it.data[it.pos-1]
}
```

## Walkthrough

For `[10, 20]`:
- pos=0 → Next() → pos=1, true. Value() → data[0] = 10.
- pos=1 → Next() → pos=2, true. Value() → data[1] = 20.
- pos=2 → Next() → false.

## Pitfalls

- Off-by-one: `Next` advances *before* the check or *after*? The pattern is
  "check, then advance" or "advance, then read pos-1".
- Calling `Value()` without `Next()` is undefined — document this.
