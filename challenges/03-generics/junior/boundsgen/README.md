# Bounds

**Level:** junior  
**Topic:** 03-generics

## Context

An axis auto-scaler needs both ends of the data range, and scanning the samples twice is wasteful.

## Task

Implement the stub(s) in [boundsgen.go](boundsgen.go):

1. Implement `Bounds`, returning the smallest element, the largest element, and `true`.
2. Return zero values and `false` for an empty slice.
3. Use a single pass over the slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Bounds([]int{3, 1, 2})
Output: 1, 3, true
```

**Example 2:**

```
Input:  Bounds([]int{5})
Output: 5, 5, true
```

**Example 3:**

```
Input:  Bounds([]int{})
Output: 0, 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Single-pass accumulation** | Two accumulators updated in one loop beat two separate scans. |
| 3 | **Zero value of `T`** | `var zero T` names the zero value of an unknown type. |

## Hint

Seed both `lo` and `hi` from `s[0]`.

## Validate

```bash
make verify
```
