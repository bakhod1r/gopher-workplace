# Memoizer

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An expensive pure function should be paid for once per input. A memoizer wraps
the function and a map: the first call for a key computes and stores, every
later call for that key just reads.

## Task

Implement `Get` on `*Memoizer` in [memorize.go](memorize.go):

1. If `key` is already in `m.cache`, return the cached value.
2. Otherwise call `m.fn(key)`, store the result under `key`, and return it.

**Constraint (senior):** across 10,000 keys asked 10 times each, the wrapped function must run exactly 10,000 times, and a cache hit must not allocate.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Get("a") with fn = k => k+"-val"
Output: "a-val"  (fn called once)
```

**Example 2:**

```
Input:  Get("a") again
Output: "a-val"  (fn not called)
```

**Example 3:**

```
Input:  Get("b") after Get("a")
Output: "b-val"  (fn called once more; caching is per key)
```

_Explanation:_ the cache is keyed, so a new key is still a miss.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok lookup** | `v, ok := m.cache[key]` is the only way to tell a miss from a cached empty string. |
| 2 | **Function stored in a struct** | `fn` is data — the memoizer works for any `func(string) string`. |
| 3 | **Pointer receiver** | Writing to the map field through a value receiver still works, but the receiver must be a pointer for the type's method set to stay consistent. |

## Hint

Do not test with `if m.cache[key] != ""`. A function that legitimately returns
`""` would then be re-called forever. Use the comma-ok form.

## Validate

```bash
make verify
```
