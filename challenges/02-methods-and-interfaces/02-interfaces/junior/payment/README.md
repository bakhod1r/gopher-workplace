# Payment Method

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A checkout charges an order through whichever payment method the customer picked.

## Task

Implement the stub(s) in [payment.go](payment.go):

1. Implement `Charge` on `Card` — succeed when the amount is at most `Limit`.
2. Implement `Charge` on `Cash` — succeed when the amount is at most `Available`.
3. Implement `Checkout`, which returns `"paid"` on success and `"declined"` otherwise.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Checkout(Card{Limit: 100}, 50)
Output: "paid"
```

**Example 2:**

```
Input:  Checkout(Card{Limit: 100}, 150)
Output: "declined"
```

**Example 3:**

```
Input:  Checkout(Cash{Available: 20}, 20)
Output: "paid"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface as a policy boundary** | Checkout logic is independent of the payment kind. |
| 2 | **Boundary conditions** | `<=` versus `<` decides the exact-amount case. |
| 3 | **Branch on a bool** | Reused: `if` returning one of two strings. |

## Hint

Exact-amount payments succeed — use `<=`.

## Validate

```bash
make verify
```
