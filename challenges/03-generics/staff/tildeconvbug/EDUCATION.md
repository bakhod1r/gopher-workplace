# The Named Type That Loses Its Width

## Intuition

`~int64` says every instantiation carries 64 bits of value. Accumulating through a 32-bit local throws away the top half of every operand, so the total wraps as soon as the data outgrows two billion.

## Approach

1. Declare the accumulator at the underlying width — `int64`.
2. Convert each element to `int64` and add it.
3. Convert the total back to `T` on the way out.

## Solution

```go
func Total[T ~int64](vals []T) T {
	var sum int64
	for _, v := range vals {
		sum += int64(v)
	}
	return T(sum)
}
```

## Walkthrough

`Total([]Millis{3000000000, 3000000000})` converts each operand to `int32`, which is already `-1294967296`, so the answer is nowhere near `6000000000`.

## Pitfalls

- Assuming a `~int64` constraint restricts the *values* as well as the underlying type.
- Accumulating into `int` — correct on 64-bit hosts, wrong on 32-bit ones.
