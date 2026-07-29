# Filter then project

## Intuition

Combine two steps in one pass: skip records failing the predicate, and collect a
projected field from the rest:

```go
for _, u := range users {
	if u.Active { out = append(out, u.Name) }
}
```

## Approach

1. Start with empty result.
2. Iterate users in order.
3. If Active, append Name.
4. Return names.

## Solution

```go
type User struct {
	Name   string
	Active bool
}

func ActiveNames(users []User) []string {
	out := []string{}
	for _, u := range users {
		if u.Active {
			out = append(out, u.Name)
		}
	}
	return out
}
```

## Walkthrough

ann active -> keep; bob inactive -> skip; cid active -> keep. [ann,cid].

## Pitfalls

- The range value is a copy — fine for reading fields.
- Order is preserved.
- Pre-size with `make([]string, 0, len(users))` if most rows pass.
