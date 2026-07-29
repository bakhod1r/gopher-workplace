# Swap Two Struct Fields

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Two struct pointers let you exchange a field between the caller's instances in
place.

## Task

Implement `SwapCounts` in [swapcounts.go](swapcounts.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a := Cart{Count:1}; b := Cart{Count:2}; SwapCounts(&a, &b)
Output: a.Count == 2, b.Count == 1
```

**Example 2:**

```
Input:  a := Cart{Count:5}; b := Cart{Count:5}; SwapCounts(&a, &b)
Output: a.Count == 5, b.Count == 5
```

**Example 3:**

```
Input:  a := Cart{Count:0}; b := Cart{Count:9}; SwapCounts(&a, &b)
Output: a.Count == 9, b.Count == 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two struct pointers** | Each aliases an instance. |
| 2 | **Field-level swap** | `a.Count, b.Count = b.Count, a.Count`. |
| 3 | **Parallel assignment** | No temporary needed. |

## Hint

`a.Count, b.Count = b.Count, a.Count`.

## Validate

```bash
make verify
```
