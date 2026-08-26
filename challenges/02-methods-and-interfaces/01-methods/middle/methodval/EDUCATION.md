# Method Values

## Intuition

In Go, `g.Greet` (without calling it) produces a **method value** — a function
that has the receiver `g` already "baked in". It's equivalent to:

```go
func(name string) string { return g.Greet(name) }
```

but shorter. The spec calls this a *bound method value*.

## Approach

1. Return `g.Greet` directly.

## Solution

```go
func ApplyMethod(g Greeter) func(string) string {
	return g.Greet
}
```

## Walkthrough

`ApplyMethod(Greeter{"Hello"})`:
- `g.Greet` creates a function bound to `Greeter{"Hello"}`.
- Calling it with `"Alice"` → `"Hello, Alice!"`.

## Pitfalls

- Wrapping in `func(name string) string { return g.Greet(name) }` works but
  is unnecessary — `g.Greet` is already the right type.
- With a value receiver, the method value captures a **copy**. If `Greeter` had
  a pointer receiver, it would capture the pointer — mutations would be visible.
