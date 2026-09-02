# Binary Search By Key

**Level:** middle  
**Topic:** 03-generics

## Context

Rows sorted by ID are searched millions of times. The rows are structs, so the comparison must go through a projection.

## Task

Implement the stub(s) in [searchbygen.go](searchbygen.go):

1. Implement `SearchBy`, returning the index of the first element whose key equals `target`.
2. When absent, return the insertion point and `false`.
3. The input is assumed sorted by `key`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SearchBy(rows, idOf, 3)
Output: index of id 3, true
```

**Example 2:**

```
Input:  SearchBy(rows, idOf, 4)
Output: insertion point, false
```

**Example 3:**

```
Input:  SearchBy(nil, idOf, 1)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Half-open search** | `lo, hi := 0, len(s)` with `hi = mid` converges on the first match. |
| 2 | **Overflow-safe midpoint** | `lo + (hi-lo)/2` never overflows, unlike `(lo+hi)/2`. |
| 3 | **Key projections** | A `func(T) K` decouples what to compare from how to traverse. |

## Hint

Converge on the lower bound, then check whether that position actually matches.

## Validate

```bash
make verify
```
