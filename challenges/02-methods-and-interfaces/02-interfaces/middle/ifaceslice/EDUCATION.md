# Interface Slices

## Intuition

`User` satisfies `Entity`, but `[]User` and `[]Entity` are different types with different element layouts — a `User` is its fields, an `Entity` element is a (type, value) pair. The conversion has to be a copy.

## Approach

1. Both `ID` methods return the stored field.
2. `ToEntities` preallocates with `make([]Entity, 0, len(users))` and appends each user.
3. `IDs` does the same shape of loop, collecting `e.ID()`.

## Solution

```go
func (u User) ID() string { return u.ID_ }

func (o Order) ID() string { return o.ID_ }

func ToEntities(users []User) []Entity {
	out := make([]Entity, 0, len(users))
	for _, u := range users {
		out = append(out, u)
	}
	return out
}

func IDs(es []Entity) []string {
	out := make([]string, 0, len(es))
	for _, e := range es {
		out = append(out, e.ID())
	}
	return out
}
```

## Walkthrough

Once converted, the `[]Entity` accepts an `Order` too — which is the point: the slice element type is the interface, so it holds any implementer.

## Pitfalls

- `return []Entity(users)` — a compile error; there is no such conversion.
- Starting from `var out []Entity` and appending, which regrows the backing array as it fills.
- Assuming the conversion is free: each element boxes into an interface header.
