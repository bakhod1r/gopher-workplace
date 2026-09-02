# Counter Interface

## Intuition

The interface is the only thing `Total` depends on. Adding a new counter kind later requires no change to `Total` at all.

## Approach

1. `(*Clicks).Count` returns `c.N`.
2. `Fixed.Count` converts the receiver: `int(f)`.
3. `Total` ranges over `cs`, adding `c.Count()` to a running sum.

## Solution

```go
func (c *Clicks) Count() int { return c.N }

func (f Fixed) Count() int { return int(f) }

func Total(cs []Counter) int {
	sum := 0
	for _, c := range cs {
		sum += c.Count()
	}
	return sum
}
```

## Walkthrough

For `[]Counter{&Clicks{N: 3}, Fixed(2)}` the loop calls `(*Clicks).Count` (3), then `Fixed.Count` (2). Sum is 5.

## Pitfalls

- `return f` in `Fixed.Count` — `Fixed` is not `int`; convert it.
- Putting `Clicks{}` (a value) in a `[]Counter` when the method has a pointer receiver — it will not compile.
- Returning early inside the loop instead of accumulating.
