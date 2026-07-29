# Boolean logic

## Intuition

`bool` has two values, `true` and `false`. The operators combine them: `&&`
(and) is true only when both sides are, `||` (or) is true when either is, `!`
negates. A "majority of three" is just every pair OR'd together:

```go
(a && b) || (a && c) || (b && c)
```

## Approach

1. A majority of three exists iff some pair is both true.
2. Return (a&&b) || (a&&c) || (b&&c).

## Solution

```go
func Majority(a, b, c bool) bool {
	return (a && b) || (a && c) || (b && c)
}
```

## Walkthrough

Majority(false,true,true): a&&b=false, a&&c=false, b&&c=true -> true.

## Pitfalls

- `&&` and `||` **short-circuit**: the right side is not evaluated if the left
  already decides the result. That matters when the right side has side effects
  or could panic (e.g. `p != nil && p.ok`).
- `!` binds tighter than `&&`, which binds tighter than `||`. Parenthesize when
  in doubt.
- Go has no implicit truthiness: only a `bool` fits where a condition is
  expected, never an int or pointer.
