# String Methods

## Intuition

Attaching a method to a struct lets you encapsulate formatting logic with the
data it formats. `person.FullName()` is cleaner than `FullName(person)`.

## Approach

1. Concatenate `First`, a space, and `Last`.

## Solution

```go
func (p Person) FullName() string {
	return p.First + " " + p.Last
}
```

## Walkthrough

For `Person{"Go", "Gopher"}`:
- `"Go" + " " + "Gopher"` → `"Go Gopher"`.

## Pitfalls

- Trimming the space when a name is empty changes the contract — the tests
  expect the space is always present.
- Using `fmt.Sprintf("%s %s", p.First, p.Last)` works but is slower for
  simple two-field concatenation.
