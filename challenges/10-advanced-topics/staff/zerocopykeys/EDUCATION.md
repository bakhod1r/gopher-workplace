# Look Up Without Copying, Store With A Copy

## Intuition

Zero-copy is not a property of the conversion but of the lifetime. A string that lives only long enough to hash and compare can safely borrow; one that the map keeps must outlive the buffer, so it has to own its bytes.

## Approach

1. Skip empty keys.
2. Build a borrowed string view for the lookup.
3. On a hit, increment through the same view.
4. On a miss, copy the bytes into a private slice and store a string over the copy.

## Solution

```go
import "unsafe"

// Count increments m's counter for each key.
//
// The lookup may borrow the key's bytes, but a key that ends up stored in
// the map must own its bytes: the caller reuses the buffers the keys point
// into.
//
// Examples:
//
// 	Count(m, [][]byte{[]byte("a")}) => m["a"] == 1
func Count(m map[string]int, keys [][]byte) {
	for _, k := range keys {
		if len(k) == 0 {
			continue
		}
		view := unsafe.String(unsafe.SliceData(k), len(k))
		if n, ok := m[view]; ok {
			m[view] = n + 1
			continue
		}
		owned := make([]byte, len(k))
		copy(owned, k)
		m[unsafe.String(unsafe.SliceData(owned), len(owned))] = 1
	}
}
```

## Walkthrough

For a key already in the map, the borrowed view hashes, matches, and is discarded — no allocation. For a new key, four bytes are copied and the stored key points at memory the caller cannot reach.

## Pitfalls

- Storing the borrowed view, which is the bug the reuse test exists to catch.
- Copying on every key, which is correct and gives up the entire optimisation.
- Reusing the borrowed view as the stored key on the miss path — it must be the copy.
