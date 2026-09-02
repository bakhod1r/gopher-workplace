# Merge Maps

**Level:** junior  
**Topic:** 03-generics

## Context

Defaults are merged with per-environment overrides at startup. Neither source map may be modified.

## Task

Implement the stub(s) in [mapscopygen.go](mapscopygen.go):

1. Implement `Merge`, returning a new map with `override` layered on top of `base`.
2. Neither input may be modified.
3. Return an empty (non-nil) map when both inputs are empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Merge({a:1}, {a:2})
Output: {a:2}
```

**Example 2:**

```
Input:  Merge({a:1}, {b:2})
Output: {a:1, b:2}
```

**Example 3:**

```
Input:  Merge(nil, nil)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`maps.Copy`** | `Copy(dst, src)` writes `src` into `dst`, overwriting collisions. |
| 2 | **Copy order decides precedence** | The later `Copy` wins, so `override` must go second. |
| 3 | **Fresh destination** | Copying into a new map is what keeps both inputs intact. |

## Hint

Two `maps.Copy` calls into a fresh map — order decides who wins.

## Validate

```bash
make verify
```
