# Field Order Is Memory

## Intuition

Each field must sit at an offset that is a multiple of its alignment. Ordering from widest to narrowest means every field lands where it already is, and only the final tail needs padding.

## Approach

1. Declare `int64`, then `int32`, then `int16`, then `bool`.
2. Return a composite literal from `NewRecord`.

## Solution

```go
type Record struct {
	ID      int64
	Count   int32
	Kind    int16
	Enabled bool
}

func NewRecord(id int64, count int32, kind int16, enabled bool) Record {
	return Record{ID: id, Count: count, Kind: kind, Enabled: enabled}
}
```

## Walkthrough

Fields occupy offsets 0, 8, 12 and 14 — fifteen bytes, padded to 16 for the `int64` alignment. Putting `Enabled` first instead pushes `ID` to offset 8 and adds seven bytes of padding before it, plus tail padding: 24 bytes.

## Pitfalls

- Assuming the compiler will reorder for you; Go deliberately does not.
- Optimising layout on a struct that exists once — it only matters in bulk.
- Hard-coding 16 as portable: on a 32-bit platform the alignment rules differ.
