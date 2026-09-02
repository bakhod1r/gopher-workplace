# Facade Pattern

## Intuition

A facade is a narrowing. Behind it there may be five types with twenty methods;
in front of it there is one call that covers the common case. It adds no
behaviour of its own — it removes choices.

## Approach

1. Hold the subsystems as fields.
2. Expose one method that calls them in the right order.
3. Combine their results into the shape the caller wants.

## Solution

```go
func (f *Facade) Operation() string {
	return f.s1.Op1() + "+" + f.s2.Op2()
}
```

## Walkthrough

`&Facade{}` gives both fields their zero values. `Sub1` and `Sub2` are empty
structs — zero bytes each — so nothing needs initializing. Go evaluates the
concatenation left to right: `"1"`, then `"+"`, then `"2"`.

## Pitfalls

- **Exposing `s1` and `s2`.** Exported fields let callers bypass the facade, and
  the ordering guarantee evaporates.
- **Confusing facade with adapter.** An adapter changes one type's shape to fit
  an expected interface; a facade hides several types behind a simpler one.
- **Putting business logic in the facade.** It should coordinate, not compute —
  otherwise it becomes the god object it was meant to prevent.
