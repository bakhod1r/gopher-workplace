# Custom JSON Unmarshaling

## Intuition

`UnmarshalJSON` must mutate the receiver, so it *must* have a pointer receiver.
The incoming `data` is raw JSON — if it's a string, it contains the literal
double-quote characters. The easiest way to strip them is to unmarshal into a
standard Go `string` first.

## Approach

1. Unmarshal `data` into a local `string`.
2. Check prefix `$` and strip it.
3. Parse dollars and cents.
4. Set `m.Cents`.

## Solution

```go
func (m *Money) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if !strings.HasPrefix(s, "$") {
		return fmt.Errorf("missing $")
	}
	var dollars, cents int
	if _, err := fmt.Sscanf(s, "$%d.%d", &dollars, &cents); err != nil {
		return err
	}
	m.Cents = dollars*100 + cents
	return nil
}
```

## Walkthrough

For `"$10.50"`:
- `data` = `[]byte("\"$10.50\"")`.
- `json.Unmarshal` into `s` gives Go string `"$10.50"`.
- `Sscanf` extracts 10 and 50.
- `m.Cents` = 1050.

## Pitfalls

- Forgetting to `json.Unmarshal(data, &s)` and trying to strip quotes manually.
- Using a value receiver — the parsed cents are lost.
- `Sscanf` can be brittle, but it's acceptable here.
