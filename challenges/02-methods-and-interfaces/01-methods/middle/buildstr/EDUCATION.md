# Builder Pattern

## Intuition

The builder pattern separates *accumulation* (`.Add()`) from *production*
(`.Build()`). Method chaining makes the API fluent and readable.

## Approach

1. Append `part` to `b.parts`.
2. Return `b`.

## Solution

```go
func (b *Builder) Add(part string) *Builder {
	b.parts = append(b.parts, part)
	return b
}
```

## Walkthrough

`NewBuilder(", ").Add("a").Add("b").Build()`:
1. `Add("a")` → parts = ["a"], returns b.
2. `Add("b")` → parts = ["a", "b"], returns b.
3. `Build()` → `strings.Join(["a", "b"], ", ")` → `"a, b"`.

## Pitfalls

- Forgetting `return b` — breaks the chain.
- Using a value receiver — `append` may reallocate, and the new slice header
  is lost on the copy.
