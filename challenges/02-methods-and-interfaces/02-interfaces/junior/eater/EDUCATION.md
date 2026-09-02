# Eater

## Intuition

An interface method can take parameters. Implementers must match the full signature — name, arguments, and results.

## Approach

1. `Cow.Eats`: `return food == "grass"`.
2. `Lion.Eats`: `return food == "meat"`.
3. `FeedableCount`: increment a counter when `e.Eats(food)` is true.

## Solution

```go
func (c Cow) Eats(food string) bool { return food == "grass" }

func (l Lion) Eats(food string) bool { return food == "meat" }

func FeedableCount(es []Eater, food string) int {
	n := 0
	for _, e := range es {
		if e.Eats(food) {
			n++
		}
	}
	return n
}
```

## Walkthrough

`FeedableCount([]Eater{Cow{}, Lion{}, Cow{}}, "grass")` calls `Cow.Eats` (true), `Lion.Eats` (false), `Cow.Eats` (true) — count 2.

## Pitfalls

- Changing the parameter name is fine; changing its type breaks satisfaction.
- Case-sensitivity: `"Grass"` is not `"grass"`.
- Returning early on the first match instead of counting all of them.
