# Split Once, Copy Nothing

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A config parser splits `key=value` lines with a helper that returns freshly allocated strings. There are thousands of lines and every one of them allocates twice for text it already has.

## Task

Implement [cutbytes.go](cutbytes.go):

1. Split `s` around the first `sep`.
2. When `sep` is absent, return `s`, `""`, false.
3. Both results must be substrings of `s` — zero allocations.

Replace the stub body in [cutbytes.go](cutbytes.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Cut("a=b", '=')
Output: "a", "b", true
```

**Example 2:**

```
Input:  Cut("a=b=c", '=')
Output: "a", "b=c", true
```

_Explanation:_ Only the first separator splits.

**Example 3:**

```
Input:  Cut("abc", '=')
Output: "abc", "", false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Substrings share bytes** | Slicing a string is a new header over the same immutable memory. |
| 2 | **First occurrence only** | The loop returns as soon as it finds one. |
| 3 | **The not-found contract** | Returning `s` unchanged lets callers ignore `found` when convenient. |

## Hint

Two slice expressions and an early return.

## Validate

```bash
make verify
```
