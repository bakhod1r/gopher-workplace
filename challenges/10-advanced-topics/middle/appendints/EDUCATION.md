# Render Numbers Without Boxing Them

## Intuition

`fmt` takes `...any`. Every argument has to become an interface value, and an interface value needs a pointer — so the int gets a heap home just to be printed. `strconv.AppendInt` writes the digits directly.

## Approach

1. Range the values, writing a space before every one but the first.
2. `dst = strconv.AppendInt(dst, int64(v), 10)`.
3. Return `dst`.

## Solution

```go
import "strconv"

// AppendInts renders vals as decimal numbers separated by ' ' and
// appends them to dst.
//
// Passing an int to a variadic any parameter puts it on the heap. Rendering
// must go straight into dst instead.
//
// Examples:
//
// 	AppendInts(nil, []int{1, 2}) => []byte("1 2")
func AppendInts(dst []byte, vals []int) []byte {
	for i, v := range vals {
		if i > 0 {
			dst = append(dst, ' ')
		}
		dst = strconv.AppendInt(dst, int64(v), 10)
	}
	return dst
}
```

## Walkthrough

Rendering 32 numbers with `fmt.Sprintf` boxes 32 ints and allocates 32 strings. The append form writes about 100 bytes into a buffer that already had room, allocating nothing.

## Pitfalls

- Forgetting to reassign `dst` — `append`'s result is the only valid slice afterwards.
- Emitting a trailing separator; the first element is the special case, not the last.
