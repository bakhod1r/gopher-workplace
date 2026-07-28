# Prefix by predicate

## The idea

Scan until the predicate fails, then slice up to that point:

```go
i := 0
for i < len(xs) && xs[i] > 0 { i++ }
return xs[:i]
```

## Why it matters

"Take while" is a streaming/parsing primitive (read digits, read whitespace,
read a header block). Unlike filter, it stops at the first failure.

## Watch out

- `xs[:i]` shares the backing array.
- Return a non-nil empty when the first element fails (`xs[:0]`).
- The dual, "drop while", slices from `i` onward.
