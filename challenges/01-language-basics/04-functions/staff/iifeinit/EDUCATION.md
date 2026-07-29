# Immediately-invoked function expressions

## Intuition

Appending `()` to a function literal runs it at once, yielding a value — handy for scoped, complex initialisation; omitting the call leaves you with the function itself.

## Approach

1. The initializer closure must be **called** to produce the map.
2. The bug assigns the function itself; return the immediately-invoked `func() map[int]int { ... }()`.

## Solution

```go
func BuildTable(n int) map[int]int {
	return func() map[int]int {
		m := map[int]int{}
		for i := 0; i < n; i++ {
			m[i] = i * i
		}
		return m
	}()
}
```

## Walkthrough

Assigning the closure to `table` leaves a function, not a map. Invoking it with `()` runs the loop and yields the populated table.

## Pitfalls

- `x := func(){...}` stores a func; `x := func(){...}()` stores its result.
- IIFEs keep setup logic local without a named helper.
