# Decorator

## Intuition

Every decorator is both an implementer and a holder of the same interface. That symmetry is what lets layers nest arbitrarily deep.

## Approach

1. `Base.Price` returns the amount.
2. Each decorator reads `Inner.Price()` once into a local, then adjusts it.
3. Multiply before dividing so integer truncation happens last.
4. `Wrap` folds the layers over `p` in order.

## Solution

```go
func (b Base) Price() int { return b.Amount }

func (d Discount) Price() int {
	p := d.Inner.Price()
	return p - p*d.Percent/100
}

func (t Tax) Price() int {
	p := t.Inner.Price()
	return p + p*t.Percent/100
}

func Wrap(p Pricer, layers ...Layer) Pricer {
	for _, l := range layers {
		p = l(p)
	}
	return p
}
```

## Walkthrough

100 discounted by 10% is 90; taxed by 20% that is `90 + 18 = 108`. The 99/33% case shows truncation: `99*33/100 = 32`, so the price is 67.

## Pitfalls

- `p * (100 - Percent) / 100` is equivalent here, but `p / 100 * Percent` truncates far too early.
- Calling `Inner.Price()` twice — correct but doubles the work of a deep chain.
- Applying `Wrap` layers back to front, which taxes before discounting.
