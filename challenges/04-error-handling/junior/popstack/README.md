# Pop From Stack

**Level:** junior
**Topic:** 04-error-handling

## Context

An undo stack hands the last action back to the editor. Popping an empty stack must not panic.

## Task

Implement `Pop` in [popstack.go](popstack.go):

1. Return the last element of `s` and the slice without it.
2. Return `nil, 0, ErrEmpty` when `s` is empty or nil.
3. Leave the original slice's remaining elements in order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Pop([]int{1, 2, 3})
Output: [1 2], 3, nil
```

**Example 2:**

```
Input:  Pop([]int{7})
Output: [], 7, nil
```

**Example 3:**

```
Input:  Pop(nil)
Output: nil, 0, ErrEmpty
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slicing** | `s[:len(s)-1]` drops the last element. |
| 2 | **Length guard** | Indexing an empty slice panics. |
| 3 | **Multiple returns** | Three results: remainder, value, error. |

## Hint

`len(s)-1` is `-1` for an empty slice — check the length before you index.

## Validate

```bash
make verify
```
