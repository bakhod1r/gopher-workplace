# Multiple assignment

## The idea

Go assigns in two phases. First it evaluates **every** expression on the right of
the `=`, then it copies those results into the operands on the left. Nothing on
the left is written until everything on the right has been read.

That is why the classic swap needs no temporary:

```go
a, b = b, a
```

The right side evaluates to the pair *(old b, old a)*. Only then does assignment
happen. In a language without this rule you would need:

```go
tmp := a
a = b
b = tmp
```

## Why it matters

The two-phase rule holds for any number of operands, which makes rotations and
shifts read as one statement:

```go
x, y, z = y, z, x     // rotate left
first, rest = rest, first
```

It also protects you from a class of bug that plagues sequential assignment:

```go
// sequential — broken
a = b   // a is now b; the old a is gone
b = a   // b gets b, not the old a
```

## Watch out

- `a, b = a, b` is legal and does nothing. Every variable is reassigned its own
  value. Easy to write by accident when the operands are longer expressions.
- The left operands must be assignable — variables, fields, indexes. You cannot
  assign to a constant or a function result.
- Declaration and assignment differ: `a, b := b, a` only compiles where `a` and
  `b` are *new* names in this scope. Inside a loop or `if`, that introduces
  shadows instead of touching the outer variables — see the *Shadowed Tally*
  puzzle.
- The count must match exactly: `a, b = f()` works when `f` returns two values,
  but `a = f()` does not compile.

## Multiple results

A Go function may return several values, and multiple assignment is how you
receive them:

```go
func Split(n, size int) (int, int)

pages, rest := Split(10, 3)   // 3, 1
```

Both must be received. If you only want one, discard the rest with the blank
identifier — see the *Discarded Remainder* puzzle.

## Try it yourself

```go
a, b := 1, 2
a, b = b, a          // 2, 1
x, y, z := 1, 2, 3
x, y, z = z, x, y    // 3, 1, 2
n, m := 5, 5
n, m = m, n          // unchanged: 5, 5
```
