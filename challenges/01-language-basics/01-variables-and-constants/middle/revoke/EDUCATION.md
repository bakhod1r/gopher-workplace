# Clearing bits with AND-NOT

## Intuition

Go has a dedicated **bit-clear** operator `&^` ("AND NOT"). `a &^ b` keeps every
bit of `a` except those set in `b`:

```go
set := Read | Write | Execute
set &^ Write // Read | Execute — Write's bit cleared, others untouched
```

It is equivalent to `a & ^b` but reads as one intent: "remove these bits".

## Approach

1. Use AND-NOT (`&^`) to clear the bits in `drop`.
2. Bits not in `drop` are untouched.

## Solution

```go
type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Execute
)

func Revoke(set, drop Permission) Permission {
	return set &^ drop
}
```

## Walkthrough

`Revoke(Read|Write, Write)` clears bit 1, leaving `Read`.

## Pitfalls

- Do not reach for `^` (XOR) to clear — that adds absent bits back.
- `&^` is left-associative and has the same precedence as `*`.
- Pairs perfectly with `|` (set bits) and `&` (test bits) for flag sets.
