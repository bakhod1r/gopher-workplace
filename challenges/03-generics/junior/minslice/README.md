# Min Of Slice

**Level:** junior  
**Topic:** 03-generics

## Context

The same monitoring panel also reports the floor of the window, using the same empty-window rule.

## Task

Implement the stub(s) in [minslice.go](minslice.go):

1. Implement `MinOf`, returning the smallest element and `true`.
2. Return the zero value and `false` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MinOf([]int{4, 1, 3})
Output: 1, true
```

**Example 2:**

```
Input:  MinOf([]string{"b", "a"})
Output: "a", true
```

**Example 3:**

```
Input:  MinOf([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Zero value of `T`** | `var zero T` names the zero value of an unknown type. |
| 3 | **Mirror of `MaxOf`** | Same seed-and-scan shape, opposite comparison. |

## Hint

Identical to `MaxOf` with `<` instead of `>`.

## Validate

```bash
make verify
```
