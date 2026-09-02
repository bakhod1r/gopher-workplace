# Last Element

**Level:** junior  
**Topic:** 03-generics

## Context

A log tail view shows the most recent entry. The buffer may not have any entries yet.

## Task

Implement the stub(s) in [lastof.go](lastof.go):

1. Implement `Last`, returning the final element and `true`.
2. For an empty slice return the zero value of `T` and `false`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Last([]int{3, 1, 4})
Output: 4, true
```

**Example 2:**

```
Input:  Last([]string{"a", "b"})
Output: "b", true
```

**Example 3:**

```
Input:  Last([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero value of `T`** | `var zero T` is the only way to name the zero value of an unknown type. |
| 2 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |
| 3 | **Index arithmetic** | Reused from language basics: the last index is `len(s)-1`. |

## Hint

`len(s)-1` on an empty slice is `-1` — check the length first.

## Validate

```bash
make verify
```
