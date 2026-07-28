# Boolean logic

## The idea

`bool` has two values, `true` and `false`. The operators combine them: `&&`
(and) is true only when both sides are, `||` (or) is true when either is, `!`
negates. A "majority of three" is just every pair OR'd together:

```go
(a && b) || (a && c) || (b && c)
```

## Why it matters

Expressing a rule directly in boolean algebra is clearer and cheaper than
converting to integers and counting. It also composes: conditions become
readable expressions instead of nested `if`s.

## Watch out

- `&&` and `||` **short-circuit**: the right side is not evaluated if the left
  already decides the result. That matters when the right side has side effects
  or could panic (e.g. `p != nil && p.ok`).
- `!` binds tighter than `&&`, which binds tighter than `||`. Parenthesize when
  in doubt.
- Go has no implicit truthiness: only a `bool` fits where a condition is
  expected, never an int or pointer.

## Try it yourself

```go
!(a && b) == (!a || !b) // De Morgan's law: always true
true || panic("x")       // won't panic — short-circuit
```
