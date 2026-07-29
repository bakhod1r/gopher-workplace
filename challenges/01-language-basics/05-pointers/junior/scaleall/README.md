# Mutate Through Slice of Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

A slice of pointers lets one call mutate many caller variables. Guard nil
entries before dereferencing.

## Task

Implement `ScaleAll` in [scaleall.go](scaleall.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a, b := 2, 3; ScaleAll([]*int{&a, &b}, 10)
Output: a == 20, b == 30
```

**Example 2:**

```
Input:  a := 5; ScaleAll([]*int{&a, nil}, 2)
Output: a == 10
```

_Explanation:_ Nil entries are skipped, no panic.

**Example 3:**

```
Input:  ScaleAll(nil, 10)
Output: no-op
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice of pointers** | Each element aliases a variable. |
| 2 | **Nil-skip** | Guard `if p != nil`. |
| 3 | **Mutate through** | `*p *= k`. |

## Hint

Range the slice; `if p != nil { *p *= k }`.

## Validate

```bash
make verify
```
