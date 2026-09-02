# Sort By Field

**Level:** junior  
**Topic:** 03-generics

## Context

A roster is displayed youngest first, and people of the same age must stay in the order the source provided.

## Task

Implement the stub(s) in [slicessortfunc.go](slicessortfunc.go):

1. Implement `SortByAge` using `slices.SortStableFunc`.
2. Compare ages with `cmp.Compare`.
3. Equal ages must keep their input order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  sort [{a 30} {b 20}]
Output: [{b 20} {a 30}]
```

**Example 2:**

```
Input:  sort [{a 20} {b 20}]
Output: [{a 20} {b 20}] (stable)
```

**Example 3:**

```
Input:  sort []
Output: no panic
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.SortFunc` vs `SortStableFunc`** | Only the stable variant preserves the order of equal elements. |
| 2 | **`cmp.Compare`** | Returns -1/0/+1 — exactly the shape the comparison function wants. |
| 3 | **Sorting non-ordered types** | Structs are not ordered, so the comparison must project a field. |

## Hint

The comparison returns an `int`, not a `bool` — `cmp.Compare(a.Age, b.Age)` gives it to you.

## Validate

```bash
make verify
```
