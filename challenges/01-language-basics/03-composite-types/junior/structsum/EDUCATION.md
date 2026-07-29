# Slices of structs

## Intuition

A struct groups named fields; a slice of structs is a table of records. Range and
read fields with dot notation:

```go
for _, o := range orders { total += o.Price * o.Qty }
```

## Approach

1. Initialize total to 0.
2. Range the orders, adding o.Price*o.Qty each iteration.
3. Return total (0 for nil/empty).

## Solution

```go
type Order struct {
	Item  string
	Price int // cents
	Qty   int
}

func Total(orders []Order) int {
	total := 0
	for _, o := range orders {
		total += o.Price * o.Qty
	}
	return total
}
```

## Walkthrough

Total(orders): pen 150*2=300; pad 300*1=300; ink 500*3=1500; total 2100.

## Pitfalls

- The range variable is a **copy** of the struct; mutating `o` doesn't change the
  slice. Use `orders[i].Field = ...` to mutate in place.
- Copying large structs each iteration has a cost; index if it matters.
- Keep money in integer cents to avoid float drift.
