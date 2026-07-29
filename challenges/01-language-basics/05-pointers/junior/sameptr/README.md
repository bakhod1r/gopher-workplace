# Pointer Identity

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Comparing pointers with `==` tests address identity, not the pointed-to values.
Two variables with equal values still have different addresses.

## Task

Implement `Same` in [sameptr.go](sameptr.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  x := 1; Same(&x, &x)
Output: true
```

_Explanation:_ Same variable → same address.

**Example 2:**

```
Input:  x, y := 1, 1; Same(&x, &y)
Output: false
```

_Explanation:_ Equal values but distinct variables → different addresses.

**Example 3:**

```
Input:  x := 1; p := &x; Same(p, &x)
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer equality** | `a == b` compares addresses. |
| 2 | **Identity vs value** | Equal values, different addresses. |
| 3 | **Distinct variables** | Each has its own address. |

## Hint

`return a == b`.

## Validate

```bash
make verify
```
