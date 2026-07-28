# Range gives you a copy

## The idea

`for _, v := range s` binds `v` to a **copy** of each element. Writing to `v`
(or its fields, for a struct) doesn't touch the slice. Mutate through the index:

```go
for i := range orders { orders[i].Price -= orders[i].Price * pct / 100 }
```

## Why it matters

"Loop and update each element" silently no-ops with the value form for value
types. It's one of the most common Go beginner-and-beyond bugs.

## Watch out

- Index (`s[i]`) to mutate; the range value is read-only in effect.
- For slices of pointers, mutating `*v` **does** work (the pointer is copied, not
  the pointee).
- Large structs copied per iteration also cost performance.
