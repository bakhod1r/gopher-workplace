# Sort By Less

**Level:** middle  
**Topic:** 03-generics

## Context

A release list is displayed in order. The ordering rule lives on the type, not in the display code.

## Task

Implement the stub(s) in [sortlessergen.go](sortlessergen.go):

1. Implement `SortedLess`, returning a stably sorted copy.
2. Derive the three-way comparison from `Less` alone.
3. Leave the input untouched; return an empty (non-nil) slice for empty or nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SortedLess([]Version{{3},{1},{2}})
Output: [{1} {2} {3}]
```

**Example 2:**

```
Input:  equal elements
Output: keep input order
```

**Example 3:**

```
Input:  SortedLess(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deriving a comparator** | Two `Less` calls yield -1, +1, or 0 — that is the whole trick. |
| 2 | **Stability needs a real zero** | Returning 0 only when neither is less is what makes ties detectable. |
| 3 | **Clone before sorting** | Sorting is destructive; the caller's slice must survive. |

## Hint

Call `Less` both ways; equal means neither direction holds.

## Validate

```bash
make verify
```
