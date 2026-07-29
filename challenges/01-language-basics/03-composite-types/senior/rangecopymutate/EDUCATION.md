# Range gives you a copy

## Intuition

`for _, v := range s` binds `v` to a **copy** of each element. Writing to `v`
(or its fields, for a struct) doesn't touch the slice. Mutate through the index:

```go
for i := range orders { orders[i].Price -= orders[i].Price * pct / 100 }
```

## Approach

1. Bug: for _, o := range orders copies each element into o; o.Price -= ... mutates the copy, so the slice is unchanged. 2. Fix: index the slice: for i := range orders { orders[i].Price -= orders[i].Price*pct/100 }. 3. Writing through orders[i] mutates the real element.

## Solution

```go
type Order struct {
	Price int
}

func Discount(orders []Order, pct int) {
	for i := range orders {
		orders[i].Price -= orders[i].Price * pct / 100
	}
}
```

## Walkthrough

o is a copy of {100}; discounting o leaves orders[0]={100}. Using orders[i], the write lands on the backing array -> {90}.

## Pitfalls

- Index (`s[i]`) to mutate; the range value is read-only in effect.
- For slices of pointers, mutating `*v` **does** work (the pointer is copied, not
  the pointee).
- Large structs copied per iteration also cost performance.
