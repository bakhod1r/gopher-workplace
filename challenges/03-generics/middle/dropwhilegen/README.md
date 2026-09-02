# Drop While

**Level:** middle  
**Topic:** 03-generics

## Context

A CSV reader skips leading blank lines and then keeps everything, blank lines included.

## Task

Implement the stub(s) in [dropwhilegen.go](dropwhilegen.go):

1. Implement `DropWhile`, skipping the leading elements that satisfy `pred`.
2. Everything from the first rejected element onward is kept, including later matches.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DropWhile([]int{2,4,5,6}, isEven)
Output: []int{5,6}
```

**Example 2:**

```
Input:  DropWhile([]int{1,2}, isEven)
Output: []int{1,2}
```

**Example 3:**

```
Input:  DropWhile([]int{2,4}, isEven)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Complement of TakeWhile** | Together they split the slice at the first failure. |
| 2 | **Index scan then copy** | Finding the split point first keeps the copy to one call. |
| 3 | **No aliasing** | Return fresh slices; sub-slices of the input share its backing array. |

## Hint

Find the split index, then copy from there.

## Validate

```bash
make verify
```
