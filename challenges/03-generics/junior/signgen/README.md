# Sign

**Level:** junior  
**Topic:** 03-generics

## Context

A trend indicator shows an arrow up, down, or flat based on a delta that may be an int count or a float rate.

## Task

Implement the stub(s) in [signgen.go](signgen.go):

1. Implement `Sign`, returning `-1`, `0`, or `1` according to the sign of `v`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Sign(-2)
Output: -1
```

**Example 2:**

```
Input:  Sign(0)
Output: 0
```

**Example 3:**

```
Input:  Sign(1.5)
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Signed sets** | Only signed types can be negative, so the set must exclude unsigned kinds. |
| 2 | **Result type is `int`** | The sign is an `int` regardless of what `T` is. |
| 3 | **Bare `switch`** | Reused from earlier: three-way results read best as a `switch`. |

## Hint

Same shape as `Compare`, with `0` as the fixed second operand.

## Validate

```bash
make verify
```
