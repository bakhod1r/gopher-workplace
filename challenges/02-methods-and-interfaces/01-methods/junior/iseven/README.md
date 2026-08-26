# Is Even

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A task scheduler distributes jobs: even-numbered tasks go to pool A, odd ones
to pool B. The check is a method on the task ID type.

## Task

Implement `IsEven` on `MyInt` in [iseven.go](iseven.go):

1. Return `true` if the integer is even (divisible by 2).
2. Return `false` otherwise.
3. Works for negative numbers too.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MyInt(4).IsEven()
Output: true
```

**Example 2:**

```
Input:  MyInt(3).IsEven()
Output: false
```

**Example 3:**

```
Input:  MyInt(0).IsEven()
Output: true
```

_Explanation:_ Zero is even.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods on defined types** | `MyInt` wraps `int` and adds behaviour. |
| 2 | **Value receiver** | Read-only check. |
| 3 | **Modulo operator** | `n % 2 == 0` tests evenness. |

## Hint

`return n%2 == 0` — modulo works with negative numbers in Go.

## Validate

```bash
make verify
```
