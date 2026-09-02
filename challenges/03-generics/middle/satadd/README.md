# Saturating Add

**Level:** middle  
**Topic:** 03-generics

## Context

A byte counter must stop at its maximum rather than wrapping to zero, which would look like a traffic collapse.

## Task

Implement the stub(s) in [satadd.go](satadd.go):

1. Implement `SatAdd`, returning the maximum value of `T` when the addition would wrap.
2. Detect the wrap without converting to a wider type.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SatAdd(uint8 style: 1, 2)
Output: 3
```

**Example 2:**

```
Input:  SatAdd(maxUint, 1)
Output: maxUint
```

**Example 3:**

```
Input:  SatAdd(uint(0), 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Wrap-around detection** | For unsigned types, `a+b < a` is true exactly when the addition wrapped. |
| 2 | **`^T(0)`** | The bitwise complement of zero is the maximum value of any unsigned type. |
| 3 | **Unsigned arithmetic is defined** | Wrapping is specified behaviour in Go, not undefined — which is why the check works. |

## Hint

`^T(0)` gives you the maximum for whatever unsigned `T` turned out to be.

## Validate

```bash
make verify
```
