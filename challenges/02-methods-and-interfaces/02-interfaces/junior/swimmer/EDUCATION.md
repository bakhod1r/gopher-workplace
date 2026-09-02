# Swimmer

## Intuition

Implementers need nothing in common structurally — only the method. An empty struct and a struct with fields sit side by side in `[]Swimmer`.

## Approach

1. `Fish.Swim` concatenates the name with `" swims"`.
2. `Duck.Swim` returns a constant.
3. `SwimAll` appends `s.Swim()` for each element.

## Solution

```go
func (f Fish) Swim() string { return f.Name + " swims" }

func (d Duck) Swim() string { return "duck swims" }

func SwimAll(ss []Swimmer) []string {
	out := make([]string, 0, len(ss))
	for _, s := range ss {
		out = append(out, s.Swim())
	}
	return out
}
```

## Walkthrough

`SwimAll` calls `Fish.Swim` (`"a swims"`), then `Duck.Swim` (`"duck swims"`), preserving the input order.

## Pitfalls

- Missing the space between the name and `swims`.
- `make([]string, len(ss))` plus `append`, which doubles the length.
- Capitalising the duck's line — the test compares exactly.
