# Unsigned underflow

## The idea

Unsigned integers have no negative range. Subtracting past zero **wraps around**
to the top of the range — modular arithmetic, not a negative result:

```go
var a uint = 2
a - 9 // 18446744073709551609 on 64-bit, not -7
```

There is no overflow panic; the wrap is silent and the value is enormous.

## Why it matters

Inventory, indices, lengths, and counters are often `uint` or `len()`-derived.
A subtraction like `have - sold` that can go negative becomes a gigantic number,
which then fails a bounds check, allocates absurd memory, or loops "forever".

## Watch out

- Guard before subtracting: `if sold >= have { return 0 }`.
- `len(x) - 1` on an empty slice underflows the same way — check for empty first.
- If you genuinely need signed differences, use a signed type and convert
  deliberately.

## Try it yourself

```go
var x uint8 = 0
x - 1 // 255 (wraps)
func sub(a, b uint) uint {
	if b > a {
		return 0
	}
	return a - b
}
```
