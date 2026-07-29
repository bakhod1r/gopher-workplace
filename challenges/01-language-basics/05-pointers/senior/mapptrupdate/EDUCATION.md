# Mutating pointees fetched from maps

## Intuition

Dereferencing into a local copies the struct; to change the stored object, write through the fetched pointer.

## Approach

1. `acc := *a` copies the struct; mutating `acc` doesn't touch the map's value.
2. Mutate through the pointer: `a.Balance += amt`.

## Solution

```go
type Account struct{ Balance int }

func Credit(m map[int]*Account, id, amt int) bool {
	a, ok := m[id]
	if !ok {
		return false
	}
	a.Balance += amt
	return true
}
```

## Walkthrough

The bug adds to a local copy that is thrown away. Writing `a.Balance += amt` updates the account the map points at.

## Pitfalls

- `acc := *a; acc.Balance += amt` edits a throwaway copy.
- `a.Balance += amt` edits the stored account.
