# Slices of structs

## The idea

A struct groups named fields; a slice of structs is a table of records. Range and
read fields with dot notation:

```go
for _, o := range orders { total += o.Price * o.Qty }
```

## Why it matters

Line items, rows, and records are modeled as structs. Summing a derived field
across them is the essence of reporting and aggregation.

## Watch out

- The range variable is a **copy** of the struct; mutating `o` doesn't change the
  slice. Use `orders[i].Field = ...` to mutate in place.
- Copying large structs each iteration has a cost; index if it matters.
- Keep money in integer cents to avoid float drift.
