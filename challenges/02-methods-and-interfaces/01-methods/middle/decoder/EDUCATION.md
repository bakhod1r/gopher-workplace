# Decoder Pattern

## Intuition

Decoder mirrors encoder: read instead of write. The `Next()/Value()` pattern
works for any sequential data source.

## Approach

1. `Next`: if `pos < len(rows)`, increment, return true.
2. `Fields`: `strings.Split(rows[pos-1], ",")`.

## Solution

```go
func (d *CSVDecoder) Next() bool {
	if d.pos >= len(d.rows) {
		return false
	}
	d.pos++
	return true
}

func (d *CSVDecoder) Fields() []string {
	return strings.Split(d.rows[d.pos-1], ",")
}
```

## Walkthrough

For `["name,age", "Alice,30"]`:
- Next() → pos=1, true. Fields() → Split("name,age") → ["name","age"].
- Next() → pos=2, true. Fields() → Split("Alice,30") → ["Alice","30"].
- Next() → false.

## Pitfalls

- Off-by-one in `pos` indexing.
- Not importing `"strings"`.
