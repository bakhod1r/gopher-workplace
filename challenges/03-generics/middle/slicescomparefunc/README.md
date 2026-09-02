# Lexicographic Comparison

**Level:** middle  
**Topic:** 03-generics

## Context

Two catalogue pages must be ordered relative to each other, the way words are ordered in a dictionary.

## Task

Implement the stub(s) in [slicescomparefunc.go](slicescomparefunc.go):

1. Implement `CompareNames` using `slices.CompareFunc`.
2. A prefix sorts before a longer slice; equal slices compare as `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CompareNames([a], [b])
Output: negative
```

**Example 2:**

```
Input:  CompareNames([a], [a,b])
Output: negative
```

**Example 3:**

```
Input:  CompareNames([a], [a])
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.CompareFunc`** | Element-wise comparison; the first difference decides. |
| 2 | **Prefix rule** | When one slice runs out first, the shorter one is smaller. |
| 3 | **Composable orderings** | The result is exactly what a sort of slices would need. |

## Hint

The first non-zero element comparison wins; length breaks the tie.

## Validate

```bash
make verify
```
