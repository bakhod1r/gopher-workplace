# Memento Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Undo needs a snapshot of past state that the outside world cannot tamper with.
A memento carries that state in an unexported field: callers can hold it and
hand it back, but only the `Editor` can read it.

## Task

Implement `Save` and `Restore` on `*Editor` in [mementopatt.go](mementopatt.go):

1. `Save()` returns a `Memento` holding the current `Text`.
2. `Restore(m)` sets `Text` back to `m.state`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Editor{Text: "initial"}.Save()
Output: Memento{state: "initial"}
```

**Example 2:**

```
Input:  save, set Text = "changed", restore
Output: Text == "initial"
```

**Example 3:**

```
Input:  restore a memento taken from an empty editor
Output: Text == ""
```

_Explanation:_ a snapshot of the zero value is still a valid snapshot.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Unexported field as encapsulation** | `state` is invisible outside the package — the memento is opaque by design. |
| 2 | **Value semantics** | `Memento` is returned by value, so the snapshot is a copy and cannot drift. |
| 3 | **Pointer receiver for mutation** | `Restore` writes to the editor, so it needs `*Editor`. |

## Hint

`Save` is `return Memento{state: e.Text}`. `Restore` is `e.Text = m.state`.
Both are one line — the interesting part is which receiver each needs.

## Validate

```bash
make verify
```
