# Slice-to-Array Length Check

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

`[4]byte(b[:4])` (slice-to-array conversion) panics when `b` has fewer than 4
bytes. A header parser must guard the length and report failure instead.

## Task

Fix the body between the markers in
[slicetoarray.go](slicetoarray.go) to check length first.

## Examples

```go
First4([]byte{1,2,3,4,5}) // => [1 2 3 4], true
First4([]byte{1,2})       // => zero, false (no panic)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice→array** | Converts a slice to a fixed array (Go 1.20+). |
| 2 | **Length requirement** | Slice must be long enough or it panics. |
| 3 | **Guard** | Check `len(b) >= 4` first. |

## Hint

`if len(b) < 4 { return [4]byte{}, false }; return [4]byte(b[:4]), true`.

## Validate

```bash
make verify
```
