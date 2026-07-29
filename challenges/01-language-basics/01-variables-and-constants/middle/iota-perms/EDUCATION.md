# iota bit flags

## Intuition

`iota` is a per-`const`-block counter: it is 0 on the first `ConstSpec` line and
increases by one on each following line. Pair it with a left shift and each flag
lands on its own bit:

```go
const (
	Read    Permission = 1 << iota // 1 << 0 == 1
	Write                          // 1 << 1 == 2
	Execute                        // 1 << 2 == 4
)
```

Only the first line carries the expression; the bare lines below **repeat** it,
and `iota` keeps climbing, so you never write the shifts by hand.

## Approach

1. Define flags with `1 << iota` so each is a distinct bit.
2. `Has` tests `set&want == want`.

## Solution

```go
type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Execute
)

func Has(set, want Permission) bool {
	return set&want == want
}
```

## Walkthrough

`Read|Write` is `0b011`; `Has(..., Read)` masks bit 0 and matches → true.

## Pitfalls

- Give the **first** constant the type (`Permission`); the whole run inherits it.
- Power-of-two spacing is what keeps flags from colliding. `1 << iota`, not
  `iota` — plain `iota` gives 0,1,2 and 0 is not a usable flag.
- A zero value means "no permissions", which is a sensible default.
