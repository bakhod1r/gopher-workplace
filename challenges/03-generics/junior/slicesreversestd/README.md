# Reverse In Place

**Level:** junior  
**Topic:** 03-generics

## Context

A timeline is rendered newest-first, but the cached slice must keep its original order for everyone else.

## Task

Implement the stub(s) in [slicesreversestd.go](slicesreversestd.go):

1. Implement `ReverseCopy`, returning a reversed copy.
2. Leave the input untouched.
3. Return an empty (non-nil) slice for an empty or nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ReverseCopy([]int{1, 2, 3})
Output: []int{3, 2, 1}
```

**Example 2:**

```
Input:  input after the call
Output: []int{1, 2, 3}
```

**Example 3:**

```
Input:  ReverseCopy([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Reverse` is in place** | It returns nothing — the argument is reordered. |
| 2 | **Cloning first** | Reused from earlier: clone before an in-place helper when the caller must keep its data. |
| 3 | **`slices.Clone` of nil** | Cloning a nil slice yields nil, so an explicit empty result needs one extra line. |

## Hint

`slices.Clone` then `slices.Reverse` — and remember `Clone(nil)` is `nil`.

## Validate

```bash
make verify
```
