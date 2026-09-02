# Payment Method

## Intuition

Checkout is written against the contract, so adding a new payment kind never touches it.

## Approach

1. Both `Charge` methods compare `amount <= <budget>`.
2. `Checkout` calls `p.Charge(amount)`.
3. Return `"paid"` when true, `"declined"` otherwise.

## Solution

```go
func (c Card) Charge(amount int) bool { return amount <= c.Limit }

func (c Cash) Charge(amount int) bool { return amount <= c.Available }

func Checkout(p Payment, amount int) string {
	if p.Charge(amount) {
		return "paid"
	}
	return "declined"
}
```

## Walkthrough

`Checkout(Card{Limit: 100}, 100)`: `100 <= 100` is true, so the result is `"paid"` — the exact-amount case the test pins down.

## Pitfalls

- Using `<` and declining exact-amount charges.
- Type-switching in `Checkout` to apply different rules — the rule belongs to the type.
- Returning a bool from `Checkout` when the signature says `string`.
