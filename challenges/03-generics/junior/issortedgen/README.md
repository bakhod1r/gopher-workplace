# Is Sorted

**Level:** junior  
**Topic:** 03-generics

## Context

A binary search is only valid on sorted input. A cheap precondition check catches misuse in tests.

## Task

Implement the stub(s) in [issortedgen.go](issortedgen.go):

1. Implement `IsSorted`, reporting whether every element is at least as large as the one before it.
2. Slices of length 0 or 1 count as sorted; equal neighbours are allowed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IsSorted([]int{1, 2, 2})
Output: true
```

**Example 2:**

```
Input:  IsSorted([]int{2, 1})
Output: false
```

**Example 3:**

```
Input:  IsSorted([]int{})
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Adjacent-pair scans** | Starting the index at 1 keeps `s[i-1]` in range. |
| 3 | **Non-decreasing vs increasing** | Equal neighbours must not fail the check. |

## Hint

Start the loop at index 1 and compare with the previous element.

## Validate

```bash
make verify
```
