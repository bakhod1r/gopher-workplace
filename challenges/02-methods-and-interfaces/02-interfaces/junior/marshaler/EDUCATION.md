# Marshaler

## Intuition

Serialisation is behaviour, so it belongs on the type. The exporter only knows `Marshal() string`.

## Approach

1. Convert both coordinates with `strconv.Itoa` and join with a comma.
2. `Label.Marshal` converts the receiver: `string(l)`.
3. `MarshalAll` appends each result into a preallocated slice.

## Solution

```go
func (p Point) Marshal() string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

func (l Label) Marshal() string { return string(l) }

func MarshalAll(ms []Marshaler) []string {
	out := make([]string, 0, len(ms))
	for _, m := range ms {
		out = append(out, m.Marshal())
	}
	return out
}
```

## Walkthrough

`Point{X: -3, Y: 0}`: `Itoa(-3)` is `"-3"`, `Itoa(0)` is `"0"`, joined as `"-3,0"`.

## Pitfalls

- `string(p.X)` instead of `strconv.Itoa(p.X)` — that converts a rune code point, not a number.
- `return l` in `Label.Marshal` — the result type is `string`, so convert.
- Forgetting the import when you introduce `strconv`.
