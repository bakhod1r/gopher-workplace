# Serialise With The Size You Can Compute

## Intuition

Every byte of the output is determined by the input, so the size is knowable before the first write. Once it is, the encoder is a straight line of appends into memory that never moves.

## Approach

1. Return nil for an empty input.
2. Sum each record's digits, one colon, and its name, plus `len(recs)-1` separators.
3. Allocate at that capacity and append the records.

## Solution

```go
import "strconv"

// Rec is one record to encode.
type Rec struct {
	ID   int
	Name string
}

// Encode renders each record as "id:name" separated by '\n'.
//
// The output's length is determined by the input, so it should be
// allocated once at exactly that size.
//
// Examples:
//
// 	Encode([]Rec{{1, "a"}}) => []byte("1:a")
func Encode(recs []Rec) []byte {
	if len(recs) == 0 {
		return nil
	}
	n := len(recs) - 1
	for _, r := range recs {
		n += len(strconv.Itoa(r.ID)) + 1 + len(r.Name)
	}
	out := make([]byte, 0, n)
	for i, r := range recs {
		if i > 0 {
			out = append(out, '\n')
		}
		out = strconv.AppendInt(out, int64(r.ID), 10)
		out = append(out, ':')
		out = append(out, r.Name...)
	}
	return out
}
```

## Walkthrough

Two records "1:a" and "22:bb" total 3 + 5 plus one separator — nine bytes, allocated once.

## Pitfalls

- Forgetting the separators, so the last append reallocates.
- Counting digits with a hand-rolled loop that mishandles negative IDs or zero.
