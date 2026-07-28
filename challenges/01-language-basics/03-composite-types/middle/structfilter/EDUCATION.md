# Filter then project

## The idea

Combine two steps in one pass: skip records failing the predicate, and collect a
projected field from the rest:

```go
for _, u := range users {
	if u.Active { out = append(out, u.Name) }
}
```

## Why it matters

"WHERE ... SELECT column" is a universal query shape. Doing it in a single range
is idiomatic and allocation-light.

## Watch out

- The range value is a copy — fine for reading fields.
- Order is preserved.
- Pre-size with `make([]string, 0, len(users))` if most rows pass.
