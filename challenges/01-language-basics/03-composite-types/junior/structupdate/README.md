# Mutate via Pointer Receiver

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

To change a struct's field from a method, the receiver must be a pointer.

## Task

Implement `Deposit(amount)` on `*Account` to add to `Balance`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=&Account{100}; a.Deposit(50)
Output: a.Balance == 150
```

**Example 2:**

```
Input:  then a.Deposit(25)
Output: a.Balance == 175
```

_Explanation:_ Pointer receiver mutates in place across calls.

**Example 3:**

```
Input:  &Account{0}; Deposit(10)
Output: Balance == 10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer receiver** | `func (a *Account)` can mutate. |
| 2 | **Value vs pointer** | Value receiver mutates a copy only. |
| 3 | **Auto-address** | `a.Deposit()` works on an addressable value. |

## Hint

`a.Balance += amount`.

## Validate

```bash
make verify
```
