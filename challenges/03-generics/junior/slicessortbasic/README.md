# Sort In Place

**Level:** junior  
**Topic:** 03-generics

## Context

A menu is rendered alphabetically. The slice belongs to the caller, who expects it sorted afterwards.

## Task

Implement the stub(s) in [slicessortbasic.go](slicessortbasic.go):

1. Implement `SortNames`, sorting the slice in place using `slices.Sort`.
2. The function returns nothing — the caller's slice must be reordered.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s := []string{"b", "a"}; SortNames(s)
Output: s == ["a", "b"]
```

**Example 2:**

```
Input:  SortNames([]string{})
Output: no panic
```

**Example 3:**

```
Input:  SortNames of one element
Output: unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Sort`** | Sorts any `~[]E` where `E` is ordered — no `less` function needed. |
| 2 | **In-place versus returning** | Some `slices` functions mutate the argument; others return a new slice. Mixing them up is the usual bug. |
| 3 | **Shared backing array** | Reused from language basics: the callee and caller see the same array. |

## Hint

`slices.Sort` mutates; there is nothing to return.

## Validate

```bash
make verify
```
