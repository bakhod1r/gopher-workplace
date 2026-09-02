# Round Trip

## Intuition

Composing `Marshaler` and `Unmarshaler` into `Codec` states the round-trip requirement in the type system: the same value can go both ways.

## Approach

1. `Marshal` joins the id and name with `|`.
2. `Unmarshal` cuts on the first `|`, parses the id, and returns `ErrBadFormat` on either failure.
3. `RoundTrip` marshals, unmarshals into `dst`, and compares `dst.Marshal()` with the original text.

## Solution

```go
func (r *Record) Marshal() string {
	return strconv.Itoa(r.ID) + "|" + r.Name
}

func (r *Record) Unmarshal(s string) error {
	idPart, name, ok := strings.Cut(s, "|")
	if !ok {
		return ErrBadFormat
	}
	id, err := strconv.Atoi(idPart)
	if err != nil {
		return ErrBadFormat
	}
	r.ID, r.Name = id, name
	return nil
}

func RoundTrip(src, dst Codec) (bool, error) {
	text := src.Marshal()
	if err := dst.Unmarshal(text); err != nil {
		return false, err
	}
	return dst.Marshal() == text, nil
}
```

## Walkthrough

`"3|a|b"` parses cleanly: `Cut` splits at the first separator, so the id is 3 and the name keeps the embedded `|`. Marshalling that record reproduces the exact input.

## Pitfalls

- `strings.Split` with an index — names containing `|` break.
- Mutating `r` before the id parse succeeds, leaving a half-filled record on error.
- Comparing structs in `RoundTrip` instead of the marshalled text, which misses formatting bugs.
