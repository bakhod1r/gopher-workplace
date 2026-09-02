# Pooled Encoder

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An encoder allocated a fresh buffer per message. At peak traffic the GC spent more time collecting buffers than the service spent encoding.

## Task

Implement the stub(s) in [poolencoder.go](poolencoder.go):

1. Implement `Encode` on `*PooledEncoder`, taking a buffer from a `sync.Pool`, using it, and returning it.
2. Implement `EncodeAll`, which encodes many messages through the `Encoder` interface.
3. Constraint: encoding must reuse buffers — the test asserts a bounded allocation count per encode, and the returned string must not alias the pooled buffer.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Encode([]string{"a", "b"})
Output: "a,b"
```

**Example 2:**

```
Input:  10000 encodes
Output: buffers are reused, allocations stay bounded
```

**Example 3:**

```
Input:  Encode(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Pool** | Reuse short-lived buffers instead of allocating per call. |
| 2 | **Aliasing hazards** | A returned string must not share memory with a recycled buffer. |
| 3 | **Reset before reuse** | A pooled buffer carries the previous call's contents. |

## Hint

`buf = buf[:0]` before use, and build the result with `string(buf)` — which copies.

## Validate

```bash
make verify
```
