# Shadowing named return values

## The idea

A short declaration inside a nested block shadows an outer variable of the same name; the outer named return keeps its zero value unless you assign to it with `=`.

## Why it matters

This is one of Go's most common real bugs — a swallowed error that `go vet`'s shadow check and linters specifically hunt for.

## Watch out

- `if x, err := f(); err != nil` shadows an outer `err`.
- Declare the extra variable separately, then assign the named return with `=`.
