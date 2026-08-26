# Overriding Promoted Methods

## Intuition

When `Pet` defines its own `String()`, it **shadows** the promoted
`Animal.String()`. Go uses the most specific (shallowest) method. But the
original is still reachable via `p.Animal.String()`.

## Approach

1. Format with `Nickname` and `Species`.

## Solution

```go
func (p Pet) String() string {
	return fmt.Sprintf("Pet(%s, %s)", p.Nickname, p.Species)
}
```

## Walkthrough

For `Pet{Animal{"Cat"}, "Whiskers"}`:
- `p.Nickname` = "Whiskers", `p.Species` = "Cat" (promoted field).
- `"Pet(Whiskers, Cat)"`.

## Pitfalls

- Calling `p.String()` inside `Pet.String()` is infinite recursion — it calls
  itself, not `Animal.String()`. Use `p.Animal.String()` explicitly if needed.
- This is not OOP overriding — it's shadowing. Go has no virtual dispatch on
  struct methods (only on interfaces).
