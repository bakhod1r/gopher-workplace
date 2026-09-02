# Any Value

## Intuition

`any` erases the static type but keeps the dynamic one. Reading it back means checking two independent things: was the key there, and is the value the type you expected.

## Approach

1. `Set` and `Len` are direct map operations.
2. `GetString` checks map presence first, then the assertion; either failure returns `"", false`.
3. `Kinds` classifies each value with a type switch into a `map[string]bool` set.
4. Collect the set's keys and `sort.Strings` them so the output is deterministic.

## Solution

```go
func (b *Bag) GetString(key string) (string, bool) {
	v, ok := b.data[key]
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	if !ok {
		return "", false
	}
	return s, true
}

func (b *Bag) Kinds() []string {
	seen := make(map[string]bool)
	for _, v := range b.data {
		switch v.(type) {
		case int:
			seen["int"] = true
		case string:
			seen["string"] = true
		case bool:
			seen["bool"] = true
		default:
			seen["other"] = true
		}
	}
	out := make([]string, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
```

## Walkthrough

Map iteration order is random in Go, so `Kinds` collects into a set and sorts before returning — otherwise the test would flake.

## Pitfalls

- Returning the map order directly — the result changes between runs.
- Skipping the presence check: a missing key yields a nil `any`, whose assertion fails anyway, but the two cases deserve separate handling for clarity.
- Using `switch v.(type)` and then referring to `v` expecting a narrowed type — bind it (`switch x := v.(type)`) when you need the value.
