# Map values are not addressable

## The idea

A map returns a **copy** of its value, and map elements aren't addressable — so
`m[key].Hits++` won't even compile. The pattern is read-modify-write:

```go
s := m[key]
s.Hits++
m[key] = s
```

## Why it matters

Struct-valued maps are common for counters and aggregates. Forgetting the
write-back silently drops every update. (Using `map[string]*Stat` avoids it, since
you mutate through the pointer.)

## Watch out

- `m[k]` yields a copy; mutating it doesn't touch the map.
- `m[k].Field = x` and `m[k].Field++` are compile errors (unaddressable).
- Pointer values (`map[K]*V`) are mutable in place.
