# Missing Or Zero

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A feature-flag lookup treats a zero count as "not configured" and silently re-enables a flag someone deliberately set to zero.

## Task

Implement [mapget.go](mapget.go):

1. Return the value under `key` and whether the key was present.
2. A stored zero must report present.
3. A nil map reports absent without panicking.

Replace the stub body in [mapget.go](mapget.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Get(map[string]int{"a":0}, "a")
Output: 0, true
```

_Explanation:_ The zero was stored.

**Example 2:**

```
Input:  Get(m, "missing")
Output: 0, false
```

**Example 3:**

```
Input:  Get(nil, "a")
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The comma-ok form** | A map index in a two-value assignment reports presence. |
| 2 | **The zero value is ambiguous** | Without the boolean, absent and zero are the same reading. |
| 3 | **Reading a nil map is legal** | It behaves as an empty map; only writing panics. |

## Hint

One statement, two results.

## Validate

```bash
make verify
```
