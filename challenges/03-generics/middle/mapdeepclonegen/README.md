# Cloning Is Shallow

**Level:** middle  
**Topic:** 03-generics

## Context

A snapshot taken with `maps.Clone` still changed under the reader, because the values were slices shared with the live map.

## Task

Implement the stub(s) in [mapdeepclonegen.go](mapdeepclonegen.go):

1. Implement `DeepClone`, copying both the map and each value slice.
2. Appending to or writing into a copied slice must not affect the original.
3. Return an empty (non-nil) map for empty or nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DeepClone({a:[1,2]})
Output: {a:[1,2]}
```

**Example 2:**

```
Input:  writing into a copied slice
Output: original unchanged
```

**Example 3:**

```
Input:  DeepClone(nil)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Shallow versus deep** | `maps.Clone` copies the map; the values keep pointing at the same arrays. |
| 2 | **One level is usually enough** | Deep-copying arbitrary nesting is a different, much harder problem. |
| 3 | **Cost is real** | The copy is O(total elements), not O(keys). |

## Hint

`maps.Clone` would share the slices — copy each value explicitly.

## Validate

```bash
make verify
```
