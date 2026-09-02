# Zip To Map

**Level:** middle  
**Topic:** 03-generics

## Context

A CSV reader has a header row and a data row and needs them married into one record.

## Task

Implement the stub(s) in [zipmapgen.go](zipmapgen.go):

1. Implement `ZipMap`, pairing `keys[i]` with `vals[i]`.
2. Stop at the shorter slice; on duplicate keys the later value wins.
3. Return an empty (non-nil) map when either slice is empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ZipMap([]string{"a","b"}, []int{1,2})
Output: {a:1, b:2}
```

**Example 2:**

```
Input:  ZipMap([]string{"a","a"}, []int{1,2})
Output: {a:2}
```

**Example 3:**

```
Input:  ZipMap([]string{"a"}, []int{})
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Positional pairing** | The minimum length bounds the loop, as in every zip. |
| 2 | **Duplicate keys collapse** | The result can be shorter than `n` — that is inherent to maps. |
| 3 | **Key constraint only** | `V` stays `any`; only keys must be comparable. |

## Hint

Take the minimum length first, then assign positionally.

## Validate

```bash
make verify
```
