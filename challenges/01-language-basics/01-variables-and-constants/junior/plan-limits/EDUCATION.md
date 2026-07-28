# Enumerations with a named type

## The idea

Go has no `enum` keyword. The idiom is a **named integer type** plus a `const`
block numbered by `iota`:

```go
type Tier int

const (
	Free Tier = iota   // 0
	Pro                // 1
	Enterprise         // 2
)
```

`Pro` and `Enterprise` repeat the previous line's expression — including its
type — so all three are `Tier`, not plain `int`.

## Why the named type earns its keep

Compare with three loose integers:

```go
const Free, Pro, Enterprise = 0, 1, 2

func Limit(t int) int
Limit(42)          // compiles: 42 is a fine int
Limit(userID)      // compiles: nothing says this is not a tier
```

With `type Tier int`, `Limit(t Tier)` rejects both. A stray `int` cannot wander
into a tier parameter; you must say `Tier(n)` and mean it. The type is the
documentation the compiler can check.

## The zero value is the first constant

`iota` starts at 0, so the first constant is also the type's zero value. That is
a design decision, not an accident: `var t Tier` gives `Free`. Choose the order
so the zero value is the safest default — the free tier, the "unset" state, the
least privileged role. If no member is a sane default, burn index 0:

```go
const (
	Unknown Tier = iota
	Free
	Pro
)
```

## Mapping a member to data

A lookup function keeps the table in one place and handles unknown input:

```go
func Limit(t Tier) int {
	switch t {
	case Pro:
		return 600
	case Enterprise:
		return 6000
	default:
		return 60      // Free, and anything unrecognised
	}
}
```

Note the `default`: a `Tier` can hold any integer — `Tier(99)` is valid Go — so
an exhaustive-looking switch still needs a fallback. The type narrows intent, it
does not constrain the value set.

## Derive related numbers

When the values follow a rule, express the rule rather than the results:

```go
const base = 60
// Pro = base * 10, Enterprise = base * 100
```

A reviewer then sees the relationship, and changing `base` moves all of them.

## Watch out

- The constants are *untyped* if you omit the type on the first line
  (`Free = iota`), which weakens the checking you wanted. Write the type once,
  on the first line; the rest inherit it.
- Ordering is an API promise. Inserting a member in the middle renumbers
  everything after it — which matters if the values are ever serialised.
- Comparing across types needs a conversion: `t == Tier(n)`, not `t == n`.

## Try it yourself

```go
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

var l Level         // Debug — the zero value
l = Error           // 3
fmt.Println(l)      // 3 — printing shows the number, not the name
```

Printing the name is a job for a method — a later topic.
