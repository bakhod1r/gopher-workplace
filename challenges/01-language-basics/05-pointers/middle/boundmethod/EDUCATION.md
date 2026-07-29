# Method values on pointer receivers

## Intuition

`c.Inc` (where Inc has a pointer receiver) yields a function value with the pointer bound, so calls mutate the shared counter.

## Approach

1. `c.Inc` is a **method value**: it binds the receiver `c` now.
2. Return it as a `func()`.
3. Each call mutates the same counter through the captured pointer.

## Solution

```go
type Counter struct{ N int }

func (c *Counter) Inc() { c.N++ }

func Bind(c *Counter) func() {
	return c.Inc
}
```

## Walkthrough

`Bind(c)` returns `c.Inc` with `c` fixed; calling the result twice increments `c.N` to 2.

## Pitfalls

- A pointer-receiver method value captures the pointer, not a copy.
- All invocations share the same underlying state.
