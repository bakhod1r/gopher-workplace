# Compare Maps

**Level:** junior  
**Topic:** 03-generics

## Context

A config reloader skips the expensive restart when the newly parsed config matches the running one.

## Task

Implement the stub(s) in [mapsequalgen.go](mapsequalgen.go):

1. Implement `SameConfig` using `maps.Equal`.
2. A nil map and an empty map count as equal.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SameConfig({a:1}, {a:1})
Output: true
```

**Example 2:**

```
Input:  SameConfig({a:1}, {a:2})
Output: false
```

**Example 3:**

```
Input:  SameConfig(nil, {})
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`maps.Equal`** | Compares sizes then every key's value — order-independent by construction. |
| 2 | **`V comparable` too** | Values must be comparable here, unlike in `maps.Clone`. |
| 3 | **The `maps` package** | Generic helpers for maps: `Clone`, `Equal`, `Copy`. |

## Hint

Both type parameters need `comparable`: the keys to index, the values to compare.

## Validate

```bash
make verify
```
