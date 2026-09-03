# Concatenating Without Regrowing

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`append(a, b...)` looks like the obvious concatenation, but it writes into `a`'s spare capacity when there is any — silently corrupting whatever else shares that array — and reallocates when there is not. Merging into a slice you sized yourself is both safe and cheap.

## Task

Implement `Merge` in [copyvsappend.go](copyvsappend.go):

1. Return the elements of `a` followed by the elements of `b`.
2. Use exactly one allocation.
3. Share no memory with either input, even after the caller appends to the result. Two empty inputs give an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  Merge([1 2], [3])
Output: [1 2 3]
```

**Example 2:**

```
Input:  Merge(nil, [3])
Output: [3]
```

**Example 3:**

```
Input:  Merge(nil, nil)
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`append(a, b...)` aliases `a`** | With spare capacity it writes in place, and the caller sees it. |
| 2 | **Size once, fill twice** | `make([]T, 0, len(a)+len(b))` needs no regrow and no second allocation. |
| 3 | **Independence is testable** | Writing to the result, and appending to it, must never reach an input. |

## Topics used again

`make` with a capacity, variadic `append`, slice aliasing.

## Hint

Allocate the exact total, then append both inputs into it.

## Validate

```bash
make verify
```
