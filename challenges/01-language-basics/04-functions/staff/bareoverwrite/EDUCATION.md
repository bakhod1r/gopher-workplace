# Bare return with deferred adjustment

## Intuition

A bare `return` returns whatever the named results currently hold; a deferred closure then transforms them. Forgetting to assign leaves the defer operating on zero.

## Approach

1. A bare `return` sends the current named return value — which was never set.
2. Assign `result = local` before returning.

## Solution

```go
func Doubled(x int) (result int) {
	defer func() { result *= 2 }()
	local := x
	_ = local
	result = local
	return
}
```

## Walkthrough

The bug returns the zero-valued `result`. Assigning the computed `local` to `result` first makes the bare return hand back 42.

## Pitfalls

- A bare `return` on an unassigned named result yields its zero value.
- The deferred edit acts on that value — assign it first.
