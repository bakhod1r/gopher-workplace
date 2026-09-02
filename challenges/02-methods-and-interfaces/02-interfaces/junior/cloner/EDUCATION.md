# Cloner

## Intuition

Returning the interface type from `Clone` lets any implementer produce its own kind. The trap is the slice: `*c` copies the header, and both configs then point at the same array.

## Approach

1. Allocate `tags := make([]string, len(c.Tags))` and `copy` into it.
2. Return `&Config{Name: c.Name, Tags: tags}`.
3. `CopyOf` returns `c.Clone()`.

## Solution

```go
func (c *Config) Clone() Cloner {
	tags := make([]string, len(c.Tags))
	copy(tags, c.Tags)
	return &Config{Name: c.Name, Tags: tags}
}

func CopyOf(c Cloner) Cloner { return c.Clone() }
```

## Walkthrough

`orig.Tags` is `["a","b"]`. `make` gives a fresh 2-element array, `copy` fills it. Writing `dup.Tags[0] = "z"` touches only the new array, so `orig.Tags[0]` is still `"a"`.

## Pitfalls

- `return &Config{Name: c.Name, Tags: c.Tags}` — shares the backing array.
- `return c` — same pointer, no clone at all.
- Forgetting that the caller must type-assert the returned `Cloner`.
