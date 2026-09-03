# Equality That Reads Too Many Fields

**Level:** staff  
**Topic:** 03-generics

## Context

A change detector compares the submitted order against the stored one and raises a revision whenever they differ. Editing the free-text note on a line now bumps the revision, and downstream systems re-price the whole order.

## Task

Fix the single planted bug in [stdequalfuncfieldbug.go](stdequalfuncfieldbug.go):

1. Find and fix the single bug so the comparison ignores the `Note` field.
2. Differences in `SKU`, `Qty`, or length must still report false.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SameLines([{x 1 "a"}], [{x 1 "b"}])
Output: true
```

**Example 2:**

```
Input:  SameLines([{x 1 "a"}], [{x 2 "a"}])
Output: false
```

**Example 3:**

```
Input:  SameLines([{x 1 ""}], [])
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Identity is a decision, not a default** | `==` on a struct compares every field, which is rarely the domain's notion of "the same". |
| 2 | **Equal versus EqualFunc** | `slices.Equal` needs `comparable` and uses `==`; `EqualFunc` lets you say what equal means. |

## Hint

`slices.Equal` compiles here. Should it?

## Validate

```bash
make verify
```
