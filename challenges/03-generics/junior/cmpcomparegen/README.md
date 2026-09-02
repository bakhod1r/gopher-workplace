# cmp.Compare

**Level:** junior  
**Topic:** 03-generics

## Context

A tag list is sorted shortest-first, with ties broken alphabetically so the output is stable across runs.

## Task

Implement the stub(s) in [cmpcomparegen.go](cmpcomparegen.go):

1. Implement `ByLengthThenName`, returning a negative number, zero, or a positive number.
2. Compare lengths first; break ties with `cmp.Compare` on the strings.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ByLengthThenName("a", "bb")
Output: negative
```

**Example 2:**

```
Input:  ByLengthThenName("bb", "aa")
Output: positive
```

**Example 3:**

```
Input:  ByLengthThenName("a", "a")
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Compare`** | The stdlib three-way comparison for ordered types. |
| 2 | **Layered comparisons** | Return early on the first non-zero result; that is how tie-breaking chains work. |
| 3 | **Comparison functions** | This is exactly the shape `slices.SortFunc` expects. |

## Hint

Return the first non-zero comparison; fall through to the tie-breaker.

## Validate

```bash
make verify
```
