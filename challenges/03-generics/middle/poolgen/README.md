# Object Pool

**Level:** middle  
**Topic:** 03-generics

## Context

A parser allocates a scratch buffer per document. Reusing them cuts allocation pressure noticeably.

## Task

Implement the stub(s) in [poolgen.go](poolgen.go):

1. Implement `NewPool`, `Get`, `Put`, and `Idle`.
2. `Get` builds a new value only when the pool is empty.
3. This pool is not safe for concurrent use — say so, do not fake it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Put(a); Get()
Output: a, no build
```

**Example 2:**

```
Input:  Get() on an empty pool
Output: a freshly built value
```

**Example 3:**

```
Input:  Put(a); Idle()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Factory functions** | `func() T` lets the pool build values of an unknown type. |
| 2 | **LIFO reuse** | Taking the most recently returned value keeps it cache-warm. |
| 3 | **Documented limits** | The stdlib `sync.Pool` adds concurrency safety and GC awareness; this one does not. |

## Hint

Take from the back — the most recently returned value is the warmest.

## Validate

```bash
make verify
```
