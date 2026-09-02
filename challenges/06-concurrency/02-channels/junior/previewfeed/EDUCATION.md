# Bounded Receives With Comma-Ok

## Intuition

`for v := range ch` runs until close — it has exactly one stop condition.
When you need two (a limit *and* a close), drop back to an explicit counted
loop with the comma-ok receive. Without the `ok` check, a closed feed would
happily pad the preview with zero ids.

## Approach

1. Start with an empty, non-nil slice.
2. Loop `i` from `0` to `limit-1`.
3. Do `id, ok := <-feed`; `break` when `!ok`.
4. Otherwise append `id`.
5. Return the slice.

## Solution

```go
func PreviewOrders(feed <-chan int, limit int) []int {
	out := []int{}
	for i := 0; i < limit; i++ {
		id, ok := <-feed
		if !ok {
			break
		}
		out = append(out, id)
	}
	return out
}
```

## Walkthrough

With ids `1, 2, 3` and `limit = 2`: two receives append `1` and `2`, the
loop counter ends it. With one id and `limit = 5`: the first receive appends
`1`, the second sees `ok == false` and breaks — `[1]`, no padding zeros.

## Pitfalls

- Dropping the `ok` check appends `0` for every receive past the close.
- `break` inside a `select` breaks the select, not the loop — here the receive is plain, so `break` exits the loop.
- A limit larger than the feed is normal, not an error.
