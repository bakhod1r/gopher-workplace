# State Machine

## Intuition

The asymmetry between reading and writing nil maps is what shapes this code: `Fire` needs no guard, while `Allow` must create the row.

## Approach

1. `Allow`: create the inner map when missing, then record the target.
2. `Fire`: look up the transition and apply it only when found.
3. `State`: return the field.

## Solution

```go
func NewMachine[S comparable, E comparable](start S) *Machine[S, E] {
	return &Machine[S, E]{state: start, table: make(map[S]map[E]S)}
}

func (m *Machine[S, E]) Allow(a S, e E, b S) {
	row, ok := m.table[a]
	if !ok {
		row = make(map[E]S)
		m.table[a] = row
	}
	row[e] = b
}

func (m *Machine[S, E]) Fire(e E) bool {
	next, ok := m.table[m.state][e]
	if !ok {
		return false
	}
	m.state = next
	return true
}

func (m *Machine[S, E]) State() S {
	return m.state
}
```

## Walkthrough

`Fire(unknown)` finds no entry, returns `false`, and leaves the machine where it was.

## Pitfalls

- Writing into a nil inner map in `Allow`, which panics.
- Assigning the state before checking whether the transition exists.
- Panicking on an unknown event instead of reporting `false`.
