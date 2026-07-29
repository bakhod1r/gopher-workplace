# iota in expressions

## Intuition

`iota` is just an integer you can put inside any constant expression. Scale or
offset it and the block repeats the expression with `iota` advancing:

```go
type Tier int
const (
	Bronze Tier = (iota + 1) * 100 // (0+1)*100 == 100
	Silver                         // (1+1)*100 == 200
	Gold                           // (2+1)*100 == 300
)
```

Only `Bronze` needs the expression; `Silver` and `Gold` inherit it.

## Approach

1. Define thresholds with `(iota + 1) * 100`.
2. `Rank` returns the highest tier whose threshold `<= score`.

## Solution

```go
type Tier int

const (
	Bronze Tier = (iota + 1) * 100
	Silver
	Gold
)

func Rank(score int) Tier {
	switch {
	case score >= int(Gold):
		return Gold
	case score >= int(Silver):
		return Silver
	case score >= int(Bronze):
		return Bronze
	default:
		return 0
	}
}
```

## Walkthrough

`Rank(250)`: 250 ≥ Silver(200) but < Gold(300), so Silver.

## Pitfalls

- `iota` starts at 0, so plain `iota * 100` gives 0,100,200. Use `(iota+1)` when
  you want the run to start at the step, not at zero.
- The expression is copied verbatim to bare lines; an explicit value on one line
  breaks the run for the lines below it.
