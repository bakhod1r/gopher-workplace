# iota bit flags

## The idea

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

## Why it matters

One byte then holds a *set* of independent permissions. Union with `|`, test
membership with `&`:

```go
set := Read | Write        // 0b011
set&Read == Read           // true
set&Execute == Execute     // false — bit not present
```

`set&want == want` is the idiom for "contains **all** of want": AND keeps only
the shared bits, and they equal `want` exactly when every wanted bit was set.

## Watch out

- Give the **first** constant the type (`Permission`); the whole run inherits it.
- Power-of-two spacing is what keeps flags from colliding. `1 << iota`, not
  `iota` — plain `iota` gives 0,1,2 and 0 is not a usable flag.
- A zero value means "no permissions", which is a sensible default.

## Try it yourself

```go
const (
	A = 1 << iota // 1
	B             // 2
	C             // 4
	D             // 8
)
```
