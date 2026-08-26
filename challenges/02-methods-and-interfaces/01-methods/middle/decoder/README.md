# Decoder

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A CSV processor reads records line by line. `Next()` advances, `Fields()`
splits the current line.

## Task

Implement `Next` and `Fields` on `*CSVDecoder` in [decoder.go](decoder.go):

1. `Next`: advance position, return whether valid.
2. `Fields`: split current row by commas.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
d := NewCSVDecoder(["name,age", "Alice,30"])
d.Next()   => true
d.Fields() => ["name", "age"]
```

**Example 2:**

```
d.Next()   => true
d.Fields() => ["Alice", "30"]
```

**Example 3:**

```
d.Next()   => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Iterator pattern** | `Next()/Fields()` mirrors `bufio.Scanner`. |
| 2 | **strings.Split** | Split a string by delimiter. |
| 3 | **Pointer receiver** | Stateful — position must persist. |

## Hint

Same pattern as the `iterator` puzzle: advance `pos`, check bounds.
`Fields` returns `strings.Split(d.rows[d.pos-1], ",")`.

## Validate

```bash
make verify
```
