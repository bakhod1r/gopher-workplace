# iota enumerations

## The idea

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

## Why it matters

The values are contiguous from 0, which means you can range over them and index
tables with them:

```go
for d := Sunday; d <= Saturday; d++ { ... }
names := [...]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
names[Tuesday] // "Tue"
```

## Watch out

- `iota` resets to 0 in **every** new `const (...)` block.
- Because the zero value is `Sunday`, an unset `Day` looks like Sunday. If 0
  should mean "invalid", start the run at `iota + 1` and reserve 0.
- Bare lines repeat the previous expression; keep them bare to stay on the
  sequence.

## Try it yourself

```go
type Suit int
const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)
```
