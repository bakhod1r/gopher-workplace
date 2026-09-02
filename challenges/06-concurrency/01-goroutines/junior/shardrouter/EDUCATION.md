# Shard Router

## Intuition

Routing is deterministic: the same key always lands on the same shard, no matter
how many goroutines are in flight. That is what makes the concurrent version
safe to compare against the sequential one.

## Approach

1. Guard `shards <= 0` on the parent goroutine.
2. Allocate `out := make([]int, len(keys))`.
3. In each goroutine fold the key's bytes into a local `h`, then write `h % shards`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package shardrouter — Gopher Workplace challenge.
package shardrouter

import (
	"sync"
)

// ShardIDs routes every key to a shard by hashing it.
//
// Examples:
//
//	ShardIDs([]string{"a", "b"}, 4)  => [1 2]
//	ShardIDs([]string{"a", "b"}, 1)  => [0 0]
//	ShardIDs([]string{"a"}, 0)       => []
func ShardIDs(keys []string, shards int) []int {
	if shards <= 0 {
		return nil
	}
	out := make([]int, len(keys))
	var wg sync.WaitGroup
	for i, key := range keys {
		wg.Add(1)
		go func(i int, key string) {
			defer wg.Done()
			h := 0
			for _, b := range []byte(key) {
				h = h*31 + int(b)
			}
			out[i] = h % shards
		}(i, key)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"a"` hashes to `97`; `97 % 4` is `1`.
- `"b"` hashes to `98`; `98 % 4` is `2`.
- With one shard everything routes to `0`, and the empty key hashes to `0`.

## Pitfalls

- Skipping the guard, so `h % shards` panics with a division by zero inside a goroutine.
- Declaring `h` outside the goroutine, which makes routing racy and wrong.
- Ranging over the key as runes instead of bytes, which changes the hash.
