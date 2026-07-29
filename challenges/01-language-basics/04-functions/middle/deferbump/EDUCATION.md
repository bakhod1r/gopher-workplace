# Defer mutating the return value

## Intuition

Because deferred functions execute after the return value is assigned, a deferred closure over a named result can transform it (wrap errors, add context).

## Approach

1. The named return `n` is set to 10.
2. A deferred `n *= 2` runs after the body, before the caller sees the value.

## Solution

```go
func Compute() (n int) {
	n = 10
	defer func() { n *= 2 }()
	return
}
```

## Walkthrough

The body sets `n = 10`; the defer doubles it to 20 during return.

## Pitfalls

- Only NAMED returns are visible to the deferred closure; a bare `return x` on an unnamed result can't be adjusted.
- The mutation happens before the caller resumes.
