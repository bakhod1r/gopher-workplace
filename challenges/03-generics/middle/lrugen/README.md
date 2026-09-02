# LRU Cache

**Level:** middle  
**Topic:** 03-generics

## Context

A read-through cache keeps the entries people actually use, not merely the ones inserted most recently.

## Task

Implement the stub(s) in [lrugen.go](lrugen.go):

1. Implement `NewLRU`, `Get`, and `Put`.
2. A successful `Get` must promote the key to most recently used.
3. `Put` evicts the least recently used entry when over capacity.
4. `touch` is provided — do not change it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  size 2: Put a,b; Get(a); Put(c)
Output: b evicted, a kept
```

**Example 2:**

```
Input:  Get(missing)
Output: zero, false
```

**Example 3:**

```
Input:  size 0: Put(a); Get(a)
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recency versus insertion** | The difference from a FIFO cache is that reads reorder the list. |
| 2 | **Two structures again** | The map answers lookups; the slice records recency. |
| 3 | **Pointer receivers mutate** | A value receiver mutates a copy, so structural changes need `*T`. |

## Hint

The only difference from a FIFO cache is that `Get` also calls `touch`.

## Validate

```bash
make verify
```
