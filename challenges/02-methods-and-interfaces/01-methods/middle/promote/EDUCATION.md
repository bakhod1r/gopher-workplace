# Struct Method Promotion

## Intuition

When you embed a struct, its methods get **promoted** to the outer type. You can
call `extended.Hello()` as if `Hello` were defined on `Extended` — Go forwards
the call to the embedded `Base`.

## Approach

1. Call `e.Hello()` — it's promoted from `Base`.
2. Concatenate with `Extra`.

## Solution

```go
func (e Extended) Describe() string {
	return e.Hello() + " [" + e.Extra + "]"
}
```

## Walkthrough

For `Extended{Base{"Go"}, "fast"}`:
- `e.Hello()` → `Base{"Go"}.Hello()` → `"Hello from Go"`.
- `+ " [fast]"` → `"Hello from Go [fast]"`.

## Pitfalls

- Thinking you need `e.Base.Hello()` — it works but `e.Hello()` is idiomatic
  (and shorter). Both compile to the same thing.
- If `Extended` also defined `Hello()`, it would **shadow** the promoted one.
