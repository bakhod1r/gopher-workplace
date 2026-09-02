# Clone

**Level:** junior  
**Topic:** 03-generics

## Context

A handler stores a slice it received from a caller who keeps writing to it. Without a copy, the stored data changes underneath.

## Task

Implement the stub(s) in [slicesclonegen.go](slicesclonegen.go):

1. Implement `Detach`, returning an independent copy of `s`.
2. Return an empty (non-nil) slice for an empty or nil input.
3. Writing to the result must not affect the input, and vice versa.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Detach([]int{1, 2})
Output: []int{1, 2}
```

**Example 2:**

```
Input:  write to the copy
Output: input unchanged
```

**Example 3:**

```
Input:  Detach(nil)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Clone`** | One allocation and one copy — replaces `make` plus `copy`. |
| 2 | **Shallow copies** | Cloning `[]*T` copies the pointers, not what they point at. |
| 3 | **Nil in, nil out** | `slices.Clone(nil)` returns nil; normalise if your API promises non-nil. |

## Hint

`slices.Clone` plus one guard for the nil case.

## Validate

```bash
make verify
```
