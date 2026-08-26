# Methods vs Functions

## Intuition

A **function** `Add(a, b int) int` lives at package scope. A **method**
`(c Calculator) Add(a, b int) int` is tied to the `Calculator` type. Methods let
you group related behaviour under a type, even if the type carries no state.

## Approach

1. Return `a + b`.

## Solution

```go
func (c Calculator) Add(a, b int) int {
	return a + b
}
```

## Walkthrough

`Calculator{}.Add(2, 3)`:
- The receiver `c` is an empty struct — unused.
- `a + b` = 5.

## Pitfalls

- An empty struct receiver seems odd, but it's a real pattern: it namespaces
  methods and can later grow fields without changing call sites.
- Do not confuse method calls (`c.Add(2,3)`) with function calls (`Add(2,3)`).
