# Deposit

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A banking service credits funds into an account. Negative deposits are rejected.

## Task

Implement `Deposit` on `*Account` in [deposit.go](deposit.go):

1. If `amount > 0`, add it to `Balance`.
2. If `amount <= 0`, do nothing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a := Account{100}; a.Deposit(50)
Output: a.Balance == 150
```

**Example 2:**

```
Input:  a := Account{100}; a.Deposit(-1)
Output: a.Balance == 100
```

**Example 3:**

```
Input:  a := Account{0}; a.Deposit(200)
Output: a.Balance == 200
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | Mutation — balance changes persist. |
| 2 | **Guard clause** | Reject invalid input early. |

## Hint

`if amount > 0 { a.Balance += amount }`.

## Validate

```bash
make verify
```
