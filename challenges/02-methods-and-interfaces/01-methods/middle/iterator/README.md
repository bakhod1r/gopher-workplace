# Iterator

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A data pipeline processes records lazily. An iterator provides `Next()` and
`Value()` methods — a common pattern in Go (e.g., `bufio.Scanner`).

## Task

Implement `Next` and `Value` on `*IntIter` in [iterator.go](iterator.go):

1. `Next` advances the position and returns `true` if there's a value.
2. `Value` returns the current element (only valid after `Next` returns true).

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
it := NewIntIter([]int{10, 20})
it.Next()  => true;  it.Value() => 10
it.Next()  => true;  it.Value() => 20
it.Next()  => false
```

**Example 2:**

```
it := NewIntIter(nil)
it.Next()  => false
```

**Example 3:**

```
it := NewIntIter([]int{42})
it.Next()  => true;  it.Value() => 42
it.Next()  => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Iterator pattern** | `Next()` advances + checks; `Value()` reads. |
| 2 | **Pointer receiver** | `pos` must persist across calls. |
| 3 | **Stateful methods** | Methods that track position via struct fields. |

## Hint

`Next`: check if `pos < len(data)`, if so the current pos is valid, increment for next call.
`Value`: return `data[pos-1]` (after Next has advanced).

## Validate

```bash
make verify
```
