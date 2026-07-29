# Fetch-then-mutate on pointer maps

## Intuition

The comma-ok read distinguishes absence; the fetched pointer allows in-place mutation of the stored object.

## Approach

1. Use comma-ok: `u, ok := m[id]`.
2. If `!ok`, return false without writing.
3. Otherwise mutate through the pointer and return true.

## Solution

```go
type User struct{ Name string }

func SetName(m map[int]*User, id int, name string) bool {
	u, ok := m[id]
	if !ok {
		return false
	}
	u.Name = name
	return true
}
```

## Walkthrough

`SetName(m, 1, "new")` finds the user, sets `u.Name`, returns true; an unknown id short-circuits to false.

## Pitfalls

- Guard the missing key before dereferencing.
- The fetched pointer aliases the stored object.
