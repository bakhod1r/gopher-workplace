# Concat From Stdlib

**Level:** middle  
**Topic:** 03-generics

## Context

Several workers return partial result slices that the coordinator must merge without disturbing any of them.

## Task

Implement the stub(s) in [slicesconcatstd.go](slicesconcatstd.go):

1. Implement `Join` using `slices.Concat`.
2. Return an empty (non-nil) slice when there is nothing to join.
3. The result must not alias any input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Join([]int{1}, []int{2,3})
Output: []int{1,2,3}
```

**Example 2:**

```
Input:  Join[int]()
Output: []int{}
```

**Example 3:**

```
Input:  writing to the result
Output: inputs unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Concat`** | Allocates once for the total length and copies every part in. |
| 2 | **Nil for no parts** | Concat of nothing is nil — normalise if your API promises non-nil. |
| 3 | **No aliasing** | The result always has its own backing array. |

## Hint

`slices.Concat` does the sizing for you; only the nil case needs a guard.

## Validate

```bash
make verify
```
