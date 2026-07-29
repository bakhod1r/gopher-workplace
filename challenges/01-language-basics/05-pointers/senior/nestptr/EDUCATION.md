# Initialising nested maps

## Intuition

Only the outer map exists after construction; each inner map is nil until assigned, and writing to nil panics.

## Approach

1. Writing `m[group][key]++` on a missing inner map assigns to a nil map → panic.
2. Lazily initialize: `if m[group] == nil { m[group] = map[string]int{} }`.

## Solution

```go
func Add(m map[string]map[string]int, group, key string) {
	if m[group] == nil {
		m[group] = map[string]int{}
	}
	m[group][key]++
}
```

## Walkthrough

The first `Add` finds no inner map and the bug panics writing to nil. Creating the inner map first makes the increment safe.

## Pitfalls

- `m[group]` is nil, not empty, until you assign a map.
- Create the inner map before incrementing.
