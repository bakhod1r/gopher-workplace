# Sizing The Map Up Front

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A map that outgrows its buckets rehashes: it allocates a bigger bucket array and migrates every key. Filling a thousand keys into a hintless map does that repeatedly. `make(map[K]V, n)` skips the whole staircase.

## Task

Implement `Index` in [preallocmap.go](preallocmap.go):

1. Map each word to the index of its *first* occurrence.
2. Create the map with a size hint derived from the input length.
3. Empty input returns an empty, non-nil map.

## Examples

**Example 1:**

```
Input:  Index([a b a c b])
Output: {a:0 b:1 c:3}
```

**Example 2:**

```
Input:  Index([a])
Output: {a:0}
```

**Example 3:**

```
Input:  Index(nil)
Output: {} (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Maps grow by rehashing** | Every growth step reinserts the existing keys into new buckets. |
| 2 | **The size hint is buckets, not a cap** | It is a preallocation, not a limit; the map still grows past it. |
| 3 | **First occurrence wins** | Only write when the key is absent, or the last index overwrites it. |

## Topics used again

Maps, the comma-ok idiom, `range` with an index.

## Hint

`make(map[string]int, len(words))`, then the comma-ok check before writing.

## Validate

```bash
make verify
```
