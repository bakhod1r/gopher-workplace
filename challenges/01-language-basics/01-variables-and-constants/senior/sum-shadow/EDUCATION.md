# Shadowing inside loops

## The idea

`:=` **declares**. Used inside a loop body or `if` block, it introduces a new
variable scoped to that block — even if a variable of the same name exists
outside. The outer one is untouched:

```go
total := 0
for _, x := range xs {
	if x > 0 {
		total := total + x // NEW total, dies at block end
		_ = total
	}
}
return total // always 0
```

Each iteration creates and discards a fresh `total`. The accumulator never grows.

## Why it matters

This compiles without warning and passes a casual read — the line *looks* like
it updates `total`. It is one of Go's most common real-world bugs, especially
with `err :=` inside loops hiding an outer `err`.

## Watch out

- Use `+=` or `=` to touch the outer variable; reserve `:=` for genuinely new
  names.
- `go vet -vettool` shadow checks and linters (e.g. `govet`'s shadow pass) catch
  many of these.
- The blank-identifier line `_ = total` here only exists to make the shadow
  compile; the real fix removes the need for it.

## Try it yourself

```go
n := 0
for i := 0; i < 3; i++ {
	n := n + 1 // shadows every iteration
	_ = n
}
// n is still 0
```
