# Slice-to-array conversion

## The idea

Go 1.20+ allows converting a slice to an array or array pointer: `[4]byte(b)`
copies the first 4 bytes. It **panics** if the slice is shorter than the array,
so guard first:

```go
if len(b) < 4 { return [4]byte{}, false }
return [4]byte(b[:4]), true
```

## Why it matters

Parsing fixed-size headers (magic numbers, lengths) benefits from array types
(comparable, value-copied). But the conversion is a runtime check — unguarded, a
truncated input crashes the parser.

## Watch out

- The slice length must be `>=` the array length, or it panics.
- `(*[4]byte)(b)` gives a pointer (no copy) with the same length rule.
- Arrays are comparable and copyable; slices are neither.
