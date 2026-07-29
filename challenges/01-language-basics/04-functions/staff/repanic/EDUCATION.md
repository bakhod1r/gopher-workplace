# Selective recover and re-panic

## Intuition

A recover that doesn't re-raise unexpected values silently swallows genuine panics; production guards absorb only known sentinels and re-panic the rest.

## Approach

1. Only the sentinel panic should be swallowed; anything else must propagate.
2. The bug recovers everything unconditionally.
3. Compare `r == ErrSentinel`; otherwise re-`panic(r)`.

## Solution

```go
const ErrSentinel = "handled"

func Run(f func()) (normal bool) {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrSentinel {
				normal = false
			} else {
				panic(r)
			}
		}
	}()
	f()
	return true
}
```

## Walkthrough

Swallowing every panic hides real failures. Checking the recovered value lets `ErrSentinel` set `normal = false` while an unexpected panic is re-raised.

## Pitfalls

- Recover then inspect `r`; re-`panic(r)` what you don't recognise.
- Absorbing everything turns real bugs into silent wrong behaviour.
