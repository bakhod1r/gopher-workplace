# Multiple assignment

## Intuition

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

## Approach

1. Return the arguments reversed: `return b, a`.

## Solution

```go
func Swap(a, b int) (int, int) {
	return b, a
}
```

## Walkthrough

Multiple return values let `Swap(1, 2)` hand back `2, 1` directly.

## Pitfalls

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
