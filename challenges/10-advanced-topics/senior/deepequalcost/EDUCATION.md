# A Comparison That Should Not Reflect

## Intuition

`DeepEqual` exists for types `==` cannot handle. Reaching for it on a comparable struct pays for a run-time type walk and two boxes to answer a question the compiler could have answered directly.

## Approach

1. Return `a != b`.

## Solution

```go
// Config is a comparable settings block.
type Config struct {
	Retries int
	Timeout int
	Name    string
	Debug   bool
}

// Changed reports whether the two configs differ.
//
// Config is a comparable struct, so == does the whole job. Reflecting over
// it boxes both operands and walks the fields at run time.
//
// Examples:
//
// 	Changed(Config{Retries: 1}, Config{Retries: 2}) => true
func Changed(a, b Config) bool {
	return a != b
}
```

## Walkthrough

`Config` holds two ints, a string and a bool — all comparable, so `!=` compares them inline. `DeepEqual` boxes both structs and walks four fields reflectively on every poll.

## Pitfalls

- Adding a slice field later, which makes the struct non-comparable and `==` a compile error — that is the signal to reconsider, not to reach back for DeepEqual.
- Comparing field by field manually, which is correct and drifts as fields are added.
