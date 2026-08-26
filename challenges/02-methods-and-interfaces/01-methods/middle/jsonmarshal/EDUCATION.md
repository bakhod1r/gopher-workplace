# Custom JSON Marshaling

## Intuition

Implementing `MarshalJSON() ([]byte, error)` on a type makes `json.Marshal`
use your method instead of the default. The output must be valid JSON — for a
string, that means including the surrounding quotes.

## Approach

1. Compute dollars and cents from `m.Cents`.
2. Format as `"$X.YY"`.
3. Return as bytes.

## Solution

```go
func (m Money) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"$%d.%02d"`, m.Cents/100, m.Cents%100)
	return []byte(s), nil
}
```

## Walkthrough

For `Money{1050}`:
- `1050 / 100` = 10, `1050 % 100` = 50.
- `"$10.50"` (with quotes for JSON string).

## Pitfalls

- Forgetting the outer quotes — JSON strings must be quoted.
- Using `%2d` instead of `%02d` — `99` cents would show as ` 99` with a space.
- `encoding/json` calls `MarshalJSON` automatically — you don't call it directly.
