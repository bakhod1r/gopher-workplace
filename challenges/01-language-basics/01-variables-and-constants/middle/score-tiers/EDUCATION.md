# iota in expressions

## The idea

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

## Why it matters

You express the *pattern* once. Change the step from 100 to 250 and every tier
updates — no hand-maintained list to drift out of sync.

## Watch out

- `iota` starts at 0, so plain `iota * 100` gives 0,100,200. Use `(iota+1)` when
  you want the run to start at the step, not at zero.
- The expression is copied verbatim to bare lines; an explicit value on one line
  breaks the run for the lines below it.

## Try it yourself

```go
const (
	_  = iota
	KB = 1 << (10 * iota)
)
const (
	Low  = iota*10 + 5 // 5
	Mid                // 15
	High               // 25
)
```
