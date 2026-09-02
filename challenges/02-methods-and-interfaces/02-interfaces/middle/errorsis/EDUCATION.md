# errors.Is

## Intuition

Wrapping adds context without destroying identity. `errors.Is` walks the `Unwrap` chain, so callers can match the sentinel however many layers were added.

## Approach

1. Use the comma-ok map read to detect the missing key.
2. Wrap with `fmt.Errorf("fetch %s: %w", key, ErrNotFound)`.
3. `IsMissing` delegates to `errors.Is`.
4. `FetchAll` returns the first error unchanged so the chain survives.

## Solution

```go
func Fetch(data map[string]string, key string) (string, error) {
	v, ok := data[key]
	if !ok {
		return "", fmt.Errorf("fetch %s: %w", key, ErrNotFound)
	}
	return v, nil
}

func IsMissing(err error) bool { return errors.Is(err, ErrNotFound) }

func FetchAll(data map[string]string, keys []string) ([]string, error) {
	out := make([]string, 0, len(keys))
	for _, k := range keys {
		v, err := Fetch(data, k)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}
```

## Walkthrough

`FetchAll(data, []string{"a", "zz"})` fails on `zz`. The error returned is the one `Fetch` built, so `errors.Is` still finds `ErrNotFound` two layers down.

## Pitfalls

- `%v` instead of `%w` — the message looks identical but `errors.Is` returns false.
- Re-wrapping the error in `FetchAll` with `%v`, which severs the chain that `Fetch` established.
- Comparing with `err == ErrNotFound`, which only matches an unwrapped sentinel.
