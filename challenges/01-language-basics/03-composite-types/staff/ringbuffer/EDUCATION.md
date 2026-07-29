# Circular indexing

## Intuition

A ring buffer stores a logical sequence in a fixed physical array. Logical index
`i` lives at physical `(head + i) mod len`:

```go
return buf[(head+i)%len(buf)]
```

The modulo is what makes the buffer circular; without it, indices run off the end.

## Approach

1. Bug: `buf[head+i]` omits the modulo, so head+i >= len(buf) indexes out of range and panics.
2. Fix: `buf[(head+i)%len(buf)]` wraps the logical index into physical bounds.

## Solution

```go
func At(buf []int, head, i int) int {
	return buf[(head+i)%len(buf)]
}
```

## Walkthrough

head=1,i=2,len=3: (1+2)%3=0 -> buf[0]=10. Without modulo, buf[3] panics.

## Pitfalls

- `%` keeps the index in `[0, len)` for non-negative operands.
- Negative logical indices need the normalize form `((x%n)+n)%n`.
- Overwriting policy (drop-oldest) is separate from the index math.
