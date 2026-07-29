# Double an Array Element

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

An array is a value; to mutate the caller's array you pass a pointer to it.
Indexing auto-dereferences.

## Task

Implement `Double` in [doubleidx.go](doubleidx.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a := [4]int{1,2,3,4}; Double(&a, 2)
Output: a[2] == 6
```

**Example 2:**

```
Input:  a := [4]int{1,2,3,4}; Double(&a, 0)
Output: a[0] == 2
```

**Example 3:**

```
Input:  a := [4]int{0,0,0,0}; Double(&a, 3)
Output: a[3] == 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer to array** | `*[4]int` aliases the caller's array. |
| 2 | **Auto-deref indexing** | `arr[i]` works on the pointer. |
| 3 | **In-place element update** | `arr[i] *= 2`. |

## Hint

`arr[i] *= 2` (Go auto-dereferences `arr`).

## Validate

```bash
make verify
```
