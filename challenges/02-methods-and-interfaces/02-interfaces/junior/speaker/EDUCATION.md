# Speaker

## Intuition

One interface variable can hold different dynamic types over its lifetime. The method that runs is chosen by whatever is stored right now.

## Approach

1. `Person.Speak` concatenates the fixed prefix with `p.Name`.
2. `Robot.Speak` returns a constant string.
3. `Introduce` returns `s.Speak()`.

## Solution

```go
func (p Person) Speak() string { return "Hi, I'm " + p.Name }

func (r Robot) Speak() string { return "I am robot" }

func Introduce(s Speaker) string { return s.Speak() }
```

## Walkthrough

Assigning `Person{Name: "Go"}` to `s` makes `s.Speak()` run `Person.Speak`. Reassigning `Robot{ID: 1}` swaps the dynamic type, and the same call now runs `Robot.Speak`.

## Pitfalls

- Including the robot's `ID` in its line — the spec says a fixed string.
- Dropping the trailing space in `"Hi, I'm "`.
- Escaping the apostrophe as `\'`, which is not valid in a Go string literal.
