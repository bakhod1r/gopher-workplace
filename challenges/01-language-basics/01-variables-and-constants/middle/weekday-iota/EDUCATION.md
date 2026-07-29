# iota enumerations

## Intuition

An enum is a small set of named integer values. `iota` numbers them for you,
starting at 0, so you never assign the numbers by hand:

```go
type Day int
const (
	Sunday Day = iota // 0
	Monday            // 1
	Tuesday           // 2
	// ...
	Saturday          // 6
)
```

Naming the type (`type Day int`) gives the constants a distinct type, so the
compiler can catch a `Day` used where a plain `int` was expected and vice versa.

## Approach

1. Define the seven days with a single `iota` run (Sunday = 0).
2. `IsWeekend` checks for Saturday or Sunday.

## Solution

```go
type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func IsWeekend(d Day) bool {
	return d == Saturday || d == Sunday
}
```

## Walkthrough

`iota` numbers the days 0–6; `IsWeekend(Saturday)` matches → true.

## Pitfalls

- `iota` resets to 0 in **every** new `const (...)` block.
- Because the zero value is `Sunday`, an unset `Day` looks like Sunday. If 0
  should mean "invalid", start the run at `iota + 1` and reserve 0.
- Bare lines repeat the previous expression; keep them bare to stay on the
  sequence.
