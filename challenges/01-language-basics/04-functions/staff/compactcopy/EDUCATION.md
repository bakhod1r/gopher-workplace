# Overlapping copy direction

## Intuition

`copy` handles overlapping ranges as if through a temporary, but the direction you choose determines whether you compact or duplicate.

## Approach

1. To drop the first element, shift the tail left: `copy(xs, xs[1:])`.
2. The bug copies right-to-left with wrong bounds, duplicating data.

## Solution

```go
func DropFirst(xs []int) []int {
	if len(xs) == 0 {
		return xs
	}
	copy(xs, xs[1:])
	return xs[:len(xs)-1]
}
```

## Walkthrough

`copy(xs[1:], xs[:len-1])` shifts the wrong way. `copy(xs, xs[1:])` moves each element down one, then reslicing drops the last stale slot.

## Pitfalls

- Drop-first shifts LEFT: `copy(xs, xs[1:])`.
- Insert shifts RIGHT: `copy(xs[i+1:], xs[i:])`.
