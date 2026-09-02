# Clone A Map

**Level:** junior  
**Topic:** 03-generics

## Context

A metrics endpoint serves a consistent view while the collector keeps writing to the live map.

## Task

Implement the stub(s) in [mapsclonegen.go](mapsclonegen.go):

1. Implement `Snapshot`, returning an independent copy of `m`.
2. Return an empty (non-nil) map for an empty or nil input.
3. Writing to the copy must not affect the original.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Snapshot({a:1})
Output: {a:1}
```

**Example 2:**

```
Input:  write to the copy
Output: original unchanged
```

**Example 3:**

```
Input:  Snapshot(nil)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `maps` package** | Generic helpers for maps: `Clone`, `Equal`, `Copy`. |
| 2 | **`maps.Clone` of nil** | Cloning a nil map returns nil — normalise when your API promises non-nil. |
| 3 | **Shallow again** | Values are copied as they are; a `map[string][]int` still shares its slices. |

## Hint

`maps.Clone` plus a nil guard.

## Validate

```bash
make verify
```
