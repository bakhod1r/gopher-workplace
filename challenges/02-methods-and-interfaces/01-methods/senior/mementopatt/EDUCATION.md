# Memento Pattern

## Intuition

Undo has three roles: the thing that changes (`Editor`), the frozen copy
(`Memento`), and whoever keeps the history (the caller). The memento is
deliberately dumb and opaque — the caller can store it in a stack but cannot
inspect or forge its contents, because `state` is unexported.

## Approach

1. `Save` copies the live field into a fresh `Memento` value.
2. `Restore` copies it back.

## Solution

```go
func (e *Editor) Save() Memento {
	return Memento{state: e.Text}
}

func (e *Editor) Restore(m Memento) {
	e.Text = m.state
}
```

## Walkthrough

`m1 := e.Save()` copies `"initial"` into `m1.state`. Assigning
`e.Text = "changed"` cannot touch `m1`, because `Memento` was returned by value
and `string` is immutable. `e.Restore(m1)` writes the saved string back, so
`e.Text` reads `"initial"` again.

## Pitfalls

- **`Save` on a pointer receiver returning `*Memento`.** Then the caller holds a
  reference into live state; if the memento ever holds a slice or map, later
  edits would corrupt the snapshot.
- **`Restore` on a value receiver.** It compiles and does nothing observable —
  the copy is discarded.
- **Exporting `state`.** The memento stops being opaque and callers start
  editing history.

## Why value semantics are enough here

`Text` is a `string`, which is immutable in Go, so a plain struct copy is a real
snapshot. Extend `Editor` with a `[]byte` or a map and `Save` would have to copy
the contents explicitly — a shallow copy would alias the live data.
