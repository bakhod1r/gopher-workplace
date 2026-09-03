# Sorted Keys Of Any String-Keyed Map

**Level:** middle
**Topic:** 10-advanced-topics / 03-reflection

## Context

A diff tool prints two configuration maps side by side. Map iteration order is randomised, so the same input produces a different diff on every run.

## Task

Implement [mapkeys.go](mapkeys.go):

1. Return the sorted keys of the string-keyed map `m`.
2. Any value type is acceptable.
3. Return nil for a non-map, a nil interface, or a map with non-string keys.

Replace the stub body in [mapkeys.go](mapkeys.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Keys(map[string]int{"b":1,"a":2})
Output: [a b]
```

**Example 2:**

```
Input:  Keys(map[int]string{1:"a"})
Output: <nil>
```

_Explanation:_ Keys must be strings.

**Example 3:**

```
Input:  Keys(map[string]int{})
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type.Key()** | The map's key type is available without touching any entry. |
| 2 | **Value.MapKeys** | Returns the keys as reflect Values, in unspecified order. |
| 3 | **Determinism** | Sorting is what makes reflective output reproducible. |

## Hint

Check the key kind through `rv.Type().Key()` before touching the entries.

## Validate

```bash
make verify
```
