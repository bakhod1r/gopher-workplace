# Block scope and shadowing

## The idea

A variable exists only inside the block (`{ ... }`) where it is declared, and
disappears at the block's closing brace. Re-declaring a name with `:=` inside a
narrower block creates a **new** variable that *shadows* the outer one:

```go
lo, hi := 0, 10
if lo > hi {
	lo, hi := hi, lo // NEW lo, hi — the outer pair is untouched!
	_ = lo; _ = hi
}
// outer lo, hi still 0, 10
```

To modify the outer variables, use `=` (assignment), not `:=` (declaration):

```go
if lo > hi {
	lo, hi = hi, lo // updates the outer pair
}
```

## Why it matters

Shadowing compiles cleanly and looks correct, which makes it a nasty class of
bug: your update evaporates when the block ends. The `=` vs `:=` distinction is
the whole difference.

## Watch out

- `:=` requires at least one new name on the left; if all names exist in the
  current scope it is a compile error — but in a *nested* block it silently
  shadows instead.
- `go vet`'s shadow analysis and linters can flag suspicious shadows.
- Multiple assignment (`lo, hi = hi, lo`) swaps without a temporary.

## Try it yourself

```go
x := 1
{
	x := 2 // shadows
	_ = x
}
// x is still 1
```
