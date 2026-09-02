# Memoize

**Level:** middle  
**Topic:** 03-generics

## Context

A pricing lookup is pure but slow. Repeated calls with the same key should cost nothing after the first.

## Task

Implement the stub(s) in [memoizegen.go](memoizegen.go):

1. Implement `Memoize`, returning a function that calls `f` at most once per distinct argument.
2. A cached zero value must not trigger a recomputation.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m := Memoize(slow); m(1); m(1)
Output: slow called once
```

**Example 2:**

```
Input:  m(1)
Output: same value both times
```

**Example 3:**

```
Input:  m := Memoize(zero); m(1); m(1)
Output: still one call
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closure-owned state** | The cache lives in the closure, one per `Memoize` call. |
| 2 | **Comma-ok again** | Only `ok` distinguishes a cached zero from a missing entry. |
| 3 | **Documented limits** | No mutex here — say so rather than pretending it is concurrency-safe. |

## Hint

Check with comma-ok: a cached `0` is a hit, not a miss.

## Validate

```bash
make verify
```
