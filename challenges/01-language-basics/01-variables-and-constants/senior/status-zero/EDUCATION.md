# The zero value and enums

## The idea

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

## Why it matters

A freshly-constructed struct, a missing map entry, or a decoded-but-absent field
all yield the zero value. If 0 is a *real* state, these look valid and skip your
"is this set?" checks.

## Watch out

- Decide deliberately what 0 means for every enum.
- `var s Status` and `Status(0)` and a struct field you forgot to set are all the
  same value.
- Stringer/validation code should treat the reserved 0 as "unknown".

## Try it yourself

```go
type Level int
const (
	_ Level = iota // 0 reserved
	Low            // 1
	High           // 2
)
```
