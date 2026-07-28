# Mode (Most Common)

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The statistical mode: the most frequent value, tie broken deterministically.

## Task

Implement `Mode(xs)`; empty → ok=false; ties → smaller value.

## Examples

```go
Mode([]int{1,2,2,3}) // => 2, true
Mode([]int{1,1,2,2}) // => 1, true (tie)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Count map** | Tally frequencies. |
| 2 | **Argmax** | Track best count and value. |
| 3 | **Deterministic ties** | Prefer the smaller value. |

## Hint

Count; scan the map tracking `(count, value)`, preferring higher count then
smaller value.

## Validate

```bash
make verify
```
