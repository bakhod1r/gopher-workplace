# Sort Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An internal sorting routine works on any collection that can report its size, compare two positions, and swap them.

## Task

Implement the stub(s) in [sortiface.go](sortiface.go):

1. Implement `Len`, `Less`, and `Swap` on `IntSlice`.
2. Implement `BubbleSort`, which sorts any `Sortable` in place using only those three methods.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IntSlice{3, 1}.Len()
Output: 2
```

**Example 2:**

```
Input:  IntSlice{3, 1}.Less(1, 0)
Output: true
```

**Example 3:**

```
Input:  s := IntSlice{3, 1, 2}; BubbleSort(s); s
Output: [1 2 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Three-method interface** | This is the shape of `sort.Interface` in the standard library. |
| 2 | **Sorting through an abstraction** | The algorithm never touches the element type. |
| 3 | **Slice mutation via index** | Reused: `s[i], s[j] = s[j], s[i]` writes through the shared backing array. |

## Hint

A slice value shares its backing array, so `Swap` on a value receiver still changes the caller's data.

## Validate

```bash
make verify
```
