# Shadowing inside loops

## Intuition

`:=` **declares**. Used inside a loop body or `if` block, it introduces a new
variable scoped to that block — even if a variable of the same name exists
outside. The outer one is untouched:

```go
total := 0
for _, x := range xs {
	if x > 0 {
		total := total + x // NEW total, dies at block end
		_ = total
	}
}
return total // always 0
```

Each iteration creates and discards a fresh `total`. The accumulator never grows.

## Approach

1. `total := total + x` declares a new inner `total`, discarded each iteration.
2. Assign to the outer one: `total = total + x`.

## Solution

```go
func SumPositive(xs []int) int {
	total := 0
	for _, x := range xs {
		if x > 0 {
			total = total + x
			_ = total
		}
	}
	return total
}
```

## Walkthrough

The `:=` shadows `total` inside the `if`, so the outer sum never grows. Using `=` accumulates correctly to 6.

## Pitfalls

- Use `+=` or `=` to touch the outer variable; reserve `:=` for genuinely new
  names.
- `go vet -vettool` shadow checks and linters (e.g. `govet`'s shadow pass) catch
  many of these.
- The blank-identifier line `_ = total` here only exists to make the shadow
  compile; the real fix removes the need for it.
