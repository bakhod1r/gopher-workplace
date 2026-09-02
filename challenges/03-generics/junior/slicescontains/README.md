# Contains From Stdlib

**Level:** junior  
**Topic:** 03-generics

## Context

The team keeps rewriting membership loops. The stdlib already ships one, tested and generic.

## Task

Implement the stub(s) in [slicescontains.go](slicescontains.go):

1. Implement `HasTag` using `slices.Contains` rather than a hand-written loop.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  HasTag([]string{"a"}, "a")
Output: true
```

**Example 2:**

```
Input:  HasTag([]int{1, 2}, 9)
Output: false
```

**Example 3:**

```
Input:  HasTag([]int{}, 1)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `slices` package** | The stdlib ships generic slice helpers — prefer them over hand-rolled loops. |
| 2 | **`slices.Contains`** | Signature: `Contains[S ~[]E, E comparable](s S, v E) bool`. |
| 3 | **Constraint matching** | Your `T comparable` lines up with what `Contains` requires. |

## Hint

One line: `return slices.Contains(tags, tag)`.

## Validate

```bash
make verify
```
