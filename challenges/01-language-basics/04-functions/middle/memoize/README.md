# Memoize a Function

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A closure can own a map as private state, turning any pure function into a
cached one without changing its signature.

## Task

Implement `Memoize` in [memoize.go](memoize.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m := Memoize(f); m(2); m(2)
Output: f called once for 2
```

**Example 2:**

```
Input:  cached result returned
Output: true
```

**Example 3:**

```
Input:  distinct inputs computed once each
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Closure-owned map** | `cache := map[int]int{}` captured privately. |
| 2 | **comma-ok lookup** | `v, ok := cache[x]` to test presence. |
| 3 | **Compute-and-store** | Miss calls `f`, stores, returns. |

## Hint

Capture `cache := map[int]int{}`; in the returned func, `if v, ok := cache[x]; ok { return v }`, else compute `f(x)`, store, return.

## Validate

```bash
make verify
```
