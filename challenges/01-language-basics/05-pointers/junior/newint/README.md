# Allocate and Return Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Returning the address of a local is safe in Go — the value escapes to the heap
and outlives the function. Each call yields a distinct pointer.

## Task

Implement `Alloc` in [newint.go](newint.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  p := Alloc(7)
Output: *p == 7
```

_Explanation:_ A fresh int holding 7 is allocated and its address returned.

**Example 2:**

```
Input:  p := Alloc(0)
Output: p != nil, *p == 0
```

**Example 3:**

```
Input:  p, q := Alloc(7), Alloc(7)
Output: p != q
```

_Explanation:_ Each call allocates a distinct int — same value, different address.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Address of a local** | `&v` is valid to return. |
| 2 | **Escape to heap** | The value outlives the call. |
| 3 | **Distinct allocations** | Each call is a new int. |

## Hint

`return &v` (v is the parameter), or `p := new(int); *p = v; return p`.

## Validate

```bash
make verify
```
