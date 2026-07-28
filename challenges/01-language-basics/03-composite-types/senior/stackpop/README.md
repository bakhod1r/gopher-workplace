# Stack Pop Shrink

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`Pop` reads the top element but returns the stack unchanged, so it never shrinks —
popping repeatedly returns the same value.

## Task

Fix the return between the markers in [stackpop.go](stackpop.go) to drop the last
element.

## Examples

```go
Pop([]int{1,2,3}) // => [1 2], 3, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice as stack** | Top is the last element. |
| 2 | **Shrink** | `s[:len(s)-1]`. |
| 3 | **Return new slice** | Caller reassigns. |

## Hint

`return s[:len(s)-1], top, true`.

## Validate

```bash
make verify
```
