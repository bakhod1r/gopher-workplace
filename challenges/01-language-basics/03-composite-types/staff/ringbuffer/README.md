# Ring Buffer Wrap

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A ring buffer maps logical index `i` to physical `(head+i) mod len`. The code
omits the modulo, so once `head+i` reaches the end it indexes out of range (and
never wraps).

## Task

Fix the index between the markers in [ringbuffer.go](ringbuffer.go) to wrap.

## Examples

```go
At([]int{10,20,30,40}, 2, 2) // => 10 (wraps)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Modular index** | `(head+i) % len(buf)`. |
| 2 | **Wraparound** | Logical past-end maps to the front. |
| 3 | **Fixed storage** | Physical size is constant. |

## Hint

`return buf[(head+i)%len(buf)]`.

## Validate

```bash
make verify
```
