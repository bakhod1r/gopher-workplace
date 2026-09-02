# Ring Buffer With Indices

**Level:** middle  
**Topic:** 03-generics

## Context

A high-rate telemetry buffer must not allocate on every sample, so the earlier append-and-trim ring is no longer good enough.

## Task

Implement the stub(s) in [ringidxgen.go](ringidxgen.go):

1. Implement `NewRing`, `Add`, and `Items`.
2. `Add` must not allocate: write into fixed storage and advance the head modulo the capacity.
3. `Items` returns a copy, oldest first.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  size 2: Add(1); Add(2); Add(3); Items()
Output: [2 3]
```

**Example 2:**

```
Input:  size 2: Items() before any Add
Output: []
```

**Example 3:**

```
Input:  size 0: Add(1); Items()
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Modular indexing** | `(i+1) % cap` wraps the head without moving any data. |
| 2 | **Fixed allocation** | The buffer is allocated once; `Add` never grows it. |
| 3 | **Deriving the oldest index** | `(head - n + cap) % cap` avoids a negative index. |

## Hint

Track a head and a count; the oldest element sits `n` positions behind the head.

## Validate

```bash
make verify
```
