# Method values on pointer receivers

## The idea

`c.Inc` (where Inc has a pointer receiver) yields a function value with the pointer bound, so calls mutate the shared counter.

## Why it matters

Bound pointer methods make convenient stateful callbacks.

## Watch out

- A pointer-receiver method value captures the pointer, not a copy.
- All invocations share the same underlying state.
