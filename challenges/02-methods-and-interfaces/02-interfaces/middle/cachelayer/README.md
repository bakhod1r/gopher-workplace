# Cache Layer

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A read-through cache sits in front of a slow source and must not call it twice for the same key.

## Task

Implement the stub(s) in [cachelayer.go](cachelayer.go):

1. Implement `Get` on `*SlowSource` — return the stored value and count the call.
2. Implement `Get` on `*Cache` — serve from memory, otherwise fetch from the wrapped source and remember the result.
3. Cache misses only: a repeated key must not reach the source again, even when the value is empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c.Get("a") twice with source {"a": "1"}
Output: "1", true both times; source.Calls == 1
```

**Example 2:**

```
Input:  c.Get("missing")
Output: "", false
```

**Example 3:**

```
Input:  a missing key fetched twice
Output: source called only once
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Read-through cache** | The cache implements the same interface it wraps, so callers cannot tell. |
| 2 | **Negative caching** | Remembering "not found" is what keeps the second lookup off the source. |
| 3 | **Comma-ok map read** | Reused: presence must be distinguishable from an empty value. |

## Hint

Cache the `(value, found)` pair, not just the value — otherwise misses fall through every time.

## Validate

```bash
make verify
```
