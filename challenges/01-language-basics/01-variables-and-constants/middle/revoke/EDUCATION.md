# Clearing bits with AND-NOT

## The idea

Go has a dedicated **bit-clear** operator `&^` ("AND NOT"). `a &^ b` keeps every
bit of `a` except those set in `b`:

```go
set := Read | Write | Execute
set &^ Write // Read | Execute — Write's bit cleared, others untouched
```

It is equivalent to `a & ^b` but reads as one intent: "remove these bits".

## Why it matters

Removing a permission is *not* subtraction and *not* XOR. Subtraction ignores
bit structure; XOR **toggles**, so clearing a bit that was already absent would
switch it on. `&^` only ever clears, which makes it idempotent:

```go
Read &^ Write // Read — dropping an absent bit changes nothing
```

## Watch out

- Do not reach for `^` (XOR) to clear — that adds absent bits back.
- `&^` is left-associative and has the same precedence as `*`.
- Pairs perfectly with `|` (set bits) and `&` (test bits) for flag sets.

## Try it yourself

```go
const (
	A = 1 << iota
	B
	C
)
all := A | B | C
all &^ B     // A | C
(A) &^ B     // A (no-op)
all &^ all   // 0
```
