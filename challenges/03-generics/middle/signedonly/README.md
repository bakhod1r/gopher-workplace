# Signed Only

**Level:** middle  
**Topic:** 03-generics

## Context

A drift monitor reports how far a reading moved and in which direction. Feeding it unsigned counters would produce nonsense.

## Task

Implement the stub(s) in [signedonly.go](signedonly.go):

1. Implement `Negate` and `AbsDiff`.
2. Both rely on `Signed`; understand what would break if `~uint` joined the set.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Negate(3)
Output: -3
```

**Example 2:**

```
Input:  AbsDiff(2, 5)
Output: 3
```

**Example 3:**

```
Input:  AbsDiff(5, 2)
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Narrow constraints on purpose** | Leaving a type out of the set is a design decision, not an oversight. |
| 2 | **Unary minus on unsigned** | `-v` on a `uint` wraps around instead of producing a negative. |
| 3 | **Subtraction underflows** | `a - b` on unsigned types with `a < b` wraps to a huge value. |

## Hint

Both bodies assume a sign exists — that assumption is what the constraint encodes.

## Validate

```bash
make verify
```
