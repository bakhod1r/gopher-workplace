# Deduplicate With Stdlib

**Level:** junior  
**Topic:** 03-generics

## Context

A tag list arrives with duplicates in arbitrary order. The rendered list is sorted and shows each tag once.

## Task

Implement the stub(s) in [dedupestdlib.go](dedupestdlib.go):

1. Implement `Distinct`, returning the distinct elements in ascending order.
2. Leave the input untouched.
3. Combine `slices.Clone`, `slices.Sort`, and `slices.Compact`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Distinct([]int{3, 1, 3})
Output: []int{1, 3}
```

**Example 2:**

```
Input:  Distinct([]string{"b", "a", "b"})
Output: []string{"a", "b"}
```

**Example 3:**

```
Input:  Distinct([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Composing stdlib helpers** | Sorting first is what turns `Compact`'s neighbour-collapse into full deduplication. |
| 2 | **Cloning first** | Reused from earlier: clone before an in-place helper when the caller must keep its data. |
| 3 | **Sorted output** | The result is ordered as a side effect — say so in the doc comment. |

## Hint

Sort first: `Compact` only collapses *adjacent* duplicates.

## Validate

```bash
make verify
```
