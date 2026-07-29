# Offsetting iota

## Intuition

`iota` counts from 0. If the first meaningful value should be the *step* (say
100), you must offset the expression, or the first constant lands on 0:

```go
const (
	Bronze Tier = iota * 100 // 0  <- a free tier!
	Silver                   // 100
	Gold                     // 200
)
```

Add one before scaling so the run starts at the step:

```go
const (
	Bronze Tier = (iota + 1) * 100 // 100
	Silver                         // 200
	Gold                           // 300
)
```

## Approach

1. `iota * 100` makes Bronze 0.
2. Offset: `(iota + 1) * 100` so Bronze is 100.

## Solution

```go
type Tier int

const (
	Bronze Tier = (iota + 1) * 100
	Silver
	Gold
)
```

## Walkthrough

The bug zeroes Bronze; `(iota + 1) * 100` gives 100, 200, 300.

## Pitfalls

- `iota * k` starts at 0; `(iota + 1) * k` starts at k.
- The offset must be inside the repeated expression so every bare line inherits
  it.
- Test the boundary values, not just the ordering.
