# Middleware Stack

## Intuition

Middleware is an onion. `Then` builds the onion from the inside out, so it must
apply the list backwards: the last middleware wraps the handler first and ends
up nearest the core, leaving the first middleware on the outside where it sees
the request earliest.

Build order and execution order are opposites. That single fact is the whole
puzzle.

## Approach

1. Let `h` be the innermost handler, `next`.
2. Walk the stack from the end, replacing `h` with `s[i](h)`.
3. Return the outermost result.

## Solution

```go
func (s Stack) Then(next Handler) Handler {
	h := next
	for i := len(s) - 1; i >= 0; i-- {
		h = s[i](h)
	}
	return h
}
```

## Walkthrough

With `Stack{mw1, mw2}`:

| step | `h` |
|------|-----|
| start | `H:req` |
| `i = 1` | `2(H:req)2` |
| `i = 0` | `1(2(H:req)2)1` |

At request time the outermost closure runs first, prints its `1(`, calls inward,
and appends `)1` on the way back out — so the string reads as a call stack.

With an empty stack the loop body never runs and `h` is still `next`, which is
the correct identity behaviour.

## Pitfalls

- **Iterating forward.** Produces `2(1(H:req)1)2`; the ordering guarantee in the
  doc comment is violated even though every middleware still runs.
- **Shadowing `next`.** Assigning to the parameter works, but a named
  accumulator makes the fold obvious.
- **A pointer receiver.** `Stack` is a slice; `Then` only reads it, so the value
  receiver is right — and it lets `Stack{mw1, mw2}.Then(h)` be written inline,
  since a composite literal is not addressable.

## Why a named slice type gets a method

Go lets any type defined in the package carry methods, including a slice type.
`type Stack []Middleware` turns "a list of middlewares" into a thing with
behaviour — `Then` — instead of a bare `[]Middleware` passed to a free function.
`sort.IntSlice` and `http.Header` are standard-library examples of the same move.
