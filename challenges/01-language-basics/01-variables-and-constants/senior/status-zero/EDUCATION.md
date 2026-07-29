# The zero value and enums

## Intuition

Every Go variable is born with its type's **zero value** — 0 for integers. For
an `iota` enum whose first member is 0, that means an unset variable silently
equals the first state:

```go
type Status int
const (
	Pending Status = iota // 0  <- also the zero value!
	Shipped               // 1
	Delivered             // 2
)
var s Status // 0 == Pending, even though nobody set it
```

Reserve 0 for "unknown/invalid" by starting the run at `iota + 1`:

```go
const (
	Pending Status = iota + 1 // 1
	Shipped                   // 2
	Delivered                 // 3
)
// zero value 0 now means "unset", distinct from every real state
```

## Approach

1. With `Pending = iota` (0), a zero-valued Status equals Pending and looks known.
2. Start at `iota + 1` so 0 is reserved for unknown.

## Solution

```go
type Status int

const (
	Pending Status = iota + 1
	Shipped
	Delivered
)

func IsKnown(s Status) bool {
	return s == Pending || s == Shipped || s == Delivered
}
```

## Walkthrough

The bug makes `Status(0)` == Pending → known. Offsetting by one keeps 0 outside the defined set.

## Pitfalls

- Decide deliberately what 0 means for every enum.
- `var s Status` and `Status(0)` and a struct field you forgot to set are all the
  same value.
- Stringer/validation code should treat the reserved 0 as "unknown".
