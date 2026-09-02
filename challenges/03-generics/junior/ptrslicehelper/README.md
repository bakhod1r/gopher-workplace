# Pointers To Elements

**Level:** junior  
**Topic:** 03-generics

## Context

An API takes optional fields as pointers. Building them from a plain slice used to produce N identical pointers.

## Task

Implement the stub(s) in [ptrslicehelper.go](ptrslicehelper.go):

1. Implement `PtrsTo`, returning one pointer per element.
2. Each pointer must address its own copy — dereferencing them must reproduce the input.
3. Return an empty (non-nil) slice for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PtrsTo([]int{1, 2}) dereferenced
Output: []int{1, 2}
```

**Example 2:**

```
Input:  the two pointers
Output: different addresses
```

**Example 3:**

```
Input:  PtrsTo([]int{})
Output: []*int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Loop variables and addresses** | Each iteration must yield a distinct variable to take the address of. |
| 2 | **Go 1.22 loop scoping** | Since Go 1.22 the range variable is per-iteration, so `&v` is already safe — the explicit copy documents the intent. |
| 3 | **Pointers to type parameters** | `*T` behaves like any pointer once instantiated. |

## Hint

Take the address of a per-iteration variable, not of a shared one.

## Validate

```bash
make verify
```
