# Offsetting iota

## The idea

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

## Why it matters

A threshold of 0 means the lowest tier is earned by everyone, including scores of
zero — usually not intended. The bug is invisible unless you check the *first*
value specifically.

## Watch out

- `iota * k` starts at 0; `(iota + 1) * k` starts at k.
- The offset must be inside the repeated expression so every bare line inherits
  it.
- Test the boundary values, not just the ordering.

## Try it yourself

```go
const (
	First = (iota + 1) * 10 // 10
	Second                  // 20
)
```
