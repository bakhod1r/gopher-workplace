# Shadowing named return values

## Intuition

A short declaration inside a nested block shadows an outer variable of the same name; the outer named return keeps its zero value unless you assign to it with `=`.

## Approach

1. The named return `err` must receive the failure.
2. The bug's `:=` inside the `if` shadows `err` with a new local, so the outer `err` stays nil.
3. Assign to the named returns directly: `n, err = strconv.Atoi(s)`.

## Solution

```go
import "strconv"

func Parse(s string) (n int, err error) {
	n, err = strconv.Atoi(s)
	return
}
```

## Walkthrough

`Parse("nope")` should report an error, but the shadowed `err` is discarded at the block's end, leaving the outer `err` nil. Assigning without `:=` writes the real result.

## Pitfalls

- `if x, err := f(); err != nil` shadows an outer `err`.
- Declare the extra variable separately, then assign the named return with `=`.
