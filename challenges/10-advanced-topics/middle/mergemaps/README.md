# Merge Into The Map You Were Given

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A settings loader merges defaults, a file and the environment. Each layer built a fresh map, so the process allocated four maps to end up with one.

## Task

Implement [mergemaps.go](mergemaps.go):

1. Copy every entry of `src` into `dst`, overwriting existing keys.
2. Return how many keys were newly added.
3. Modify `dst` in place; a nil `dst` adds nothing and must not panic.
4. `src` must not be modified.

Replace the stub body in [mergemaps.go](mergemaps.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Merge({"a":1}, {"a":9,"b":2})
Output: 1, dst is {"a":9,"b":2}
```

_Explanation:_ Only "b" is new.

**Example 2:**

```
Input:  Merge(dst, nil)
Output: 0, dst unchanged
```

**Example 3:**

```
Input:  Merge(nil, src)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Maps are reference-like** | Writing through the parameter reaches the caller's map. |
| 2 | **Comma-ok for presence** | Distinguishes a missing key from a key holding zero. |
| 3 | **In-place over rebuilding** | The destination already has its buckets. |

## Hint

Check presence before writing, or the count is wrong.

## Validate

```bash
make verify
```
