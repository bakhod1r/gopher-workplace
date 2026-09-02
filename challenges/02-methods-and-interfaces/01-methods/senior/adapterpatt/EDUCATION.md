# Adapter Pattern

## Intuition

Two types do the same job but speak different languages. You cannot edit the
legacy one (it is someone else's package, or frozen by compatibility). So you
introduce a third type whose only work is translation: it holds the old type,
and its methods present the new shape.

The pattern lives entirely in the method set. `ModernAdapter` has one method,
`GetIntData`, and inside it the string world ends and the int world begins.

## Approach

1. Ask the wrapped value for what it can give: a `string`.
2. Convert at the boundary with `strconv.Atoi`.
3. Decide what an unconvertible value means. Here the contract says `0`.

## Solution

```go
func (a *ModernAdapter) GetIntData() int {
	n, err := strconv.Atoi(a.legacy.GetStringData())
	if err != nil {
		return 0
	}
	return n
}
```

## Walkthrough

`a.legacy` is a `LegacySystem` value — the zero value works, since the type has
no fields. `GetStringData` has a value receiver, so calling it through the
pointer `a` is fine: Go dereferences automatically.

`strconv.Atoi("123")` yields `(123, nil)`, and the adapter returns `123`. On
bad input it yields `(0, *strconv.NumError)`, and the guard returns the
explicit `0`.

## Pitfalls

- **Returning `n` without checking `err`.** It happens to be `0` today, but
  relying on that couples you to an undocumented detail.
- **Adapting on a value receiver.** `func (a ModernAdapter)` would also compile,
  but the pointer receiver keeps the method set consistent if the adapter later
  gains cached state.
- **Doing the conversion at every call site instead.** That is the bug the
  adapter exists to prevent.
