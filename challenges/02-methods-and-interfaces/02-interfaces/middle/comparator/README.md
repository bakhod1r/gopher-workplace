# Comparator

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A report sorts the same records by different keys, chosen at run time.

## Task

Implement the stub(s) in [comparator.go](comparator.go):

1. Implement `Compare` on `ByName` and `ByAge` (ascending).
2. Implement `Compare` on `Reverse`, which inverts the wrapped comparator.
3. Implement `SortWith`, which sorts a copy of the records with the given comparator and leaves the input untouched.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SortWith(recs, ByAge{})
Output: youngest first
```

**Example 2:**

```
Input:  SortWith(recs, Reverse{Inner: ByAge{}})
Output: oldest first
```

**Example 3:**

```
Input:  the input slice after SortWith
Output: unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comparator as a value** | The ordering policy is chosen at run time, not compile time. |
| 2 | **Wrapping a comparator** | `Reverse` decorates any comparator, including itself. |
| 3 | **Copy before sort** | Efficiency and safety: one allocation, no surprise mutation of the caller's slice. |

## Hint

`sort.SliceStable` with a closure calling `c.Compare(...) < 0` keeps the order deterministic.

## Validate

```bash
make verify
```
