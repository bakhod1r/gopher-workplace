# Deleting map entries

## Intuition

Assigning a nil (or zero) value keeps the key; the `delete` built-in is the only way to remove an entry from a map.

## Approach

1. `m[id] = nil` keeps the key with a nil value — length unchanged.
2. `delete(m, id)` removes the entry entirely.

## Solution

```go
func Remove(m map[int]*int, id int) {
	delete(m, id)
}
```

## Walkthrough

Assigning nil leaves the key present, so `len(m)` and `ok` still report it. `delete` actually removes it.

## Pitfalls

- `m[id] = nil` leaves the key present with a nil value.
- `delete(m, id)` removes it entirely.
