# Deferred call ordering

## Intuition

Each `defer` pushes onto a stack unwound in reverse at function exit; deferred closures see and can modify named return values.

## Approach

1. Deferred calls run in **last-in, first-out** order.
2. Defer appends 1, 2, 3; they fire 3, 2, 1.

## Solution

```go
func Order() (out []int) {
	defer func() { out = append(out, 1) }()
	defer func() { out = append(out, 2) }()
	defer func() { out = append(out, 3) }()
	return
}
```

## Walkthrough

The three defers queue in order but execute in reverse, so `out` ends `[3 2 1]`.

## Pitfalls

- Schedule order is source order; execution order is reversed.
- Append 1,2,3 in source → result [3 2 1].
