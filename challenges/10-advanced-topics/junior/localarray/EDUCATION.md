# A Scratch Array That Never Leaves

## Intuition

A local array whose address never leaves the function is just part of the frame. That makes it free — no allocator, no collector, no cost beyond the stack pointer already moving.

## Approach

1. Declare a fixed local array big enough for any int.
2. Take the absolute value.
3. Peel digits with `%10` and `/10`, counting until the number reaches zero.

## Solution

```go
// Digits returns the number of decimal digits in the absolute value of n,
// computing them into a fixed-size local array.
//
// The array never escapes, so the whole function is allocation-free.
//
// Examples:
//
// 	Digits(1234) => 4
func Digits(n int) int {
	var buf [20]byte
	if n < 0 {
		n = -n
	}
	i := 0
	for {
		buf[i] = byte('0' + n%10)
		i++
		n /= 10
		if n == 0 {
			return i
		}
	}
}
```

## Walkthrough

1234 peels to 4, 3, 2, 1 — four iterations, so the answer is 4. The array is written but never read and never escapes, so the frame is the only memory involved.

## Pitfalls

- Returning a slice of the local array — that makes it escape.
- Looping `for n > 0`, which returns 0 for the input 0.
