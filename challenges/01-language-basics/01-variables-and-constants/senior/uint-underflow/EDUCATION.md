# Unsigned underflow

## Intuition

Unsigned integers have no negative range. Subtracting past zero **wraps around**
to the top of the range — modular arithmetic, not a negative result:

```go
var a uint = 2
a - 9 // 18446744073709551609 on 64-bit, not -7
```

There is no overflow panic; the wrap is silent and the value is enormous.

## Approach

1. `have - sold` on `uint` wraps when `sold > have`.
2. Guard: `if sold >= have { return 0 }` before subtracting.

## Solution

```go
func Remaining(have, sold uint) uint {
	if sold >= have {
		return 0
	}
	return have - sold
}
```

## Walkthrough

`Remaining(2, 9)` would wrap to a huge value; the guard returns 0 instead.

## Pitfalls

- Guard before subtracting: `if sold >= have { return 0 }`.
- `len(x) - 1` on an empty slice underflows the same way — check for empty first.
- If you genuinely need signed differences, use a signed type and convert
  deliberately.
