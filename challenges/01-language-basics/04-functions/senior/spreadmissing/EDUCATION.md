# Spreading a slice into a variadic call

## The idea

`f(xs...)` forwards each element of `xs`; without the `...` you pass a single value (here the length), not the contents.

## Why it matters

Forgetting the spread is a common bug when wrapping variadic APIs.

## Watch out

- `sum(xs...)` spreads; `sum(xs)` won't compile for `...int`, and `sum(len(xs))` compiles but is wrong.
- A nil slice spreads to zero arguments.
