# Read Through A Pointer That May Be Nil

## Intuition

A pointer type has one more state than its base type: absent. The dereference is only defined for the other states, so the check is not defensive noise — it is the missing case.

## Approach

1. Return 0 when `p` is nil.
2. Otherwise return `*p`.

## Solution

```go
// Value returns what p points at, or 0 when p is nil.
//
// A nil pointer is a valid value; dereferencing one is not.
//
// Examples:
//
// 	Value(nil) => 0
func Value(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
```

## Walkthrough

`Value(&n)` reads the current contents of `n`, so a write to `n` between taking the address and calling is visible. `Value(nil)` returns the documented default instead of faulting.

## Pitfalls

- Returning `*p` after a `p != nil` check written the wrong way round.
- Treating 0 as "absent" at the call site — `&zero` is present and zero.
