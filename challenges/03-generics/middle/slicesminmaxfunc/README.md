# Min And Max By Comparison

**Level:** middle  
**Topic:** 03-generics

## Context

A catalogue view highlights the cheapest and most expensive item, and the filtered list is often empty.

## Task

Implement the stub(s) in [slicesminmaxfunc.go](slicesminmaxfunc.go):

1. Implement `Cheapest` and `Priciest` using `slices.MinFunc` and `slices.MaxFunc`.
2. Both panic on an empty slice — guard first and report `false`.
3. Ties resolve to the first such element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Cheapest(items)
Output: the lowest-priced item, true
```

**Example 2:**

```
Input:  Priciest(items)
Output: the highest-priced item, true
```

**Example 3:**

```
Input:  Cheapest(nil)
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`MinFunc` and `MaxFunc`** | Comparison-based extremes for types with no natural ordering. |
| 2 | **Documented panics** | Both panic on an empty slice — the wrapper's whole job is the guard. |
| 3 | **First-match ties** | `MinFunc` returns the first minimal element, which keeps output stable. |

## Hint

Guard the empty slice, then one stdlib call each.

## Validate

```bash
make verify
```
