# Unmarshaler

## Intuition

Parsing mutates the receiver, so `Unmarshal` needs a pointer. Returning an error keeps malformed input out of the domain type.

## Approach

1. `strings.Cut(s, "=")` gives key, value, and whether the separator was found.
2. Return `ErrBadPair` when it is not found.
3. Otherwise assign both fields and return `nil`.
4. `UnmarshalAll` parses into a fresh `Pair` per line and returns early on error.

## Solution

```go
func (p *Pair) Unmarshal(s string) error {
	k, v, ok := strings.Cut(s, "=")
	if !ok {
		return ErrBadPair
	}
	p.Key, p.Value = k, v
	return nil
}

func UnmarshalAll(lines []string) ([]Pair, error) {
	out := make([]Pair, 0, len(lines))
	for _, line := range lines {
		var p Pair
		if err := p.Unmarshal(line); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}
```

## Walkthrough

`"k=v=w"`: `Cut` splits at the *first* `=`, so the key is `k` and the value keeps the rest, `"v=w"`. `"k="` is valid with an empty value.

## Pitfalls

- `strings.Split` and taking index 1 — that loses everything after a second `=`.
- Treating an empty value as an error; only a missing `=` is malformed.
- Reusing one `Pair` variable across iterations without clearing it, so a failed parse leaks old fields.
