# Variadic Methods

## Intuition

Variadic methods accept a flexible number of arguments. `Encode(fields ...string)`
can be called with any number of strings — including zero.

## Approach

1. `strings.Join(fields, ",")`.
2. `e.Rows = append(e.Rows, row)`.

## Solution

```go
func (e *CSVEncoder) Encode(fields ...string) {
	e.Rows = append(e.Rows, strings.Join(fields, ","))
}
```

(Add `"strings"` to imports.)

## Walkthrough

`e.Encode("a", "b", "c")`:
- `strings.Join(["a","b","c"], ",")` → `"a,b,c"`.
- `e.Rows = append(e.Rows, "a,b,c")`.

## Pitfalls

- `fmt.Sprint` doesn't join with commas — it adds spaces.
- Need to import `"strings"` for `Join`.
