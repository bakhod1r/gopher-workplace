# Binary Min-Heap

**Level:** middle  
**Topic:** 03-generics

## Context

A scheduler always needs the earliest deadline next. Sorting the whole queue on every insert is too slow.

## Task

Implement the stub(s) in [heapgen.go](heapgen.go):

1. Implement `Push`, `Pop`, and `Len` for a binary min-heap.
2. `Pop` returns the smallest element; the zero value and `false` when empty.
3. Keep the heap invariant: every parent is `<=` its children.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(3); Push(1); Pop()
Output: 1, true
```

**Example 2:**

```
Input:  Pop() on an empty heap
Output: 0, false
```

**Example 3:**

```
Input:  Push(2); Push(1); Len()
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Structural invariants** | A data structure is defined by what stays true after every operation. |
| 2 | **Array-encoded tree** | Children of `i` live at `2i+1` and `2i+2`; the parent at `(i-1)/2`. |
| 3 | **Sift up and down** | Push repairs upward from the new leaf; Pop repairs downward from the root. |

## Hint

Push sifts up from the last index; Pop moves the last element to the root and sifts down.

## Validate

```bash
make verify
```
