# Unique By Key

**Level:** middle  
**Topic:** 03-generics

## Context

An import feed repeats records with the same ID but different timestamps. The first sighting wins.

## Task

Implement the stub(s) in [uniquebygen.go](uniquebygen.go):

1. Implement `UniqueBy`, keeping the first element for each distinct key.
2. Preserve input order; the elements themselves need not be comparable.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  UniqueBy(users, idOf)
Output: one user per id, first wins
```

**Example 2:**

```
Input:  UniqueBy([]int{1,11,2}, mod10)
Output: []int{1,2}
```

**Example 3:**

```
Input:  UniqueBy([]int{}, mod10)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Key projections** | A `func(T) K` decouples what to compare from how to traverse. |
| 2 | **Only the key is comparable** | `T any` lets struct elements through; only `K` needs `comparable`. |
| 3 | **Stable order** | Appending in traversal order preserves the input's relative order. |

## Hint

The `seen` set is keyed by `K`, not by `T`.

## Validate

```bash
make verify
```
