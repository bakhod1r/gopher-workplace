# Pointer vs Value Receiver Choice

## Intuition

If a method modifies the receiver, it **must** be a pointer receiver. If it only
reads, either works but value is conventional for small structs.

Here `Inc` mutates, so it needs `*Counter`. `Value` only reads, so it uses
`Counter`.

## Approach

1. Increment `c.N` by 1.

## Solution

```go
func (c *Counter) Inc() {
	c.N++
}
```

## Walkthrough

For `Counter{0}` → `Inc()`:
- `c.N++` → `c.N` = 1.

## Pitfalls

- Using a value receiver `(c Counter)` — `c.N++` increments the copy, caller
  sees no change.
- Mixing pointer and value receivers on the same type is allowed but confusing;
  here it's intentional to teach the difference.
