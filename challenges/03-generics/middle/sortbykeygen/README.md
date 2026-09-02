# Sort By Key

**Level:** middle  
**Topic:** 03-generics

## Context

A table is sorted by whichever column the user clicked, and rows with equal values must not jump around.

## Task

Implement the stub(s) in [sortbykeygen.go](sortbykeygen.go):

1. Implement `SortedBy`, returning a stably sorted copy ordered by `key`.
2. Leave the input untouched; return an empty (non-nil) slice for empty or nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SortedBy(people, ageOf)
Output: youngest first
```

**Example 2:**

```
Input:  equal keys
Output: keep input order
```

**Example 3:**

```
Input:  SortedBy(nil, ageOf)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.SortStableFunc`** | Stability is what stops equal rows from shuffling between renders. |
| 2 | **`cmp.Compare` on the projection** | Returns the -1/0/+1 the sort wants. |
| 3 | **Key projections** | A `func(T) K` decouples what to compare from how to traverse. |

## Hint

Clone first, then `slices.SortStableFunc` with `cmp.Compare(key(a), key(b))`.

## Validate

```bash
make verify
```
