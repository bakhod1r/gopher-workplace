# Count Distinct

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A set is a map whose keys are the members. Counting distinct values is inserting
each into a set and taking its size.

## Task

Implement `Distinct(xs)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  []int{1,2,2,3,3,3}
Output: 3
```

**Example 2:**

```
Input:  []int{5,5,5}
Output: 1
```

**Example 3:**

```
Input:  nil
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Set as map** | `map[int]struct{}` keys are members. |
| 2 | **Idempotent insert** | Re-inserting a key is a no-op. |
| 3 | **Size = len** | `len(set)` is the distinct count. |

## Hint

`seen := make(map[int]struct{}); for _, x := range xs { seen[x] = struct{}{} }; return len(seen)`.

## Validate

```bash
make verify
```
