# Method On A Slice Type

**Level:** junior  
**Topic:** 03-generics

## Context

A pipeline reads better as `data.Filter(ok).Len()` than as nested function calls.

## Task

Implement the stub(s) in [slicetypegen.go](slicetypegen.go):

1. Implement `Filter`, returning the accepted elements as a `Slice[T]`.
2. Implement `Len`.
3. Preserve order and leave the receiver unmodified.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Slice[int]{1,2,3}.Filter(isEven)
Output: Slice[int]{2}
```

**Example 2:**

```
Input:  Slice[int]{1,2,3}.Filter(isEven).Len()
Output: 1
```

**Example 3:**

```
Input:  Slice[int]{}.Filter(isEven)
Output: Slice[int]{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generic named slice types** | `type Slice[T any] []T` can carry methods — the underlying type is still a slice. |
| 2 | **Chaining** | Returning `Slice[T]` rather than `[]T` is what allows `.Filter(...).Len()`. |
| 3 | **Value receivers on slices** | The receiver is a slice header copy; not reusing it keeps the caller's data intact. |

## Hint

`make(Slice[T], 0, len(s))` — the named type works with `make` exactly like `[]T`.

## Validate

```bash
make verify
```
