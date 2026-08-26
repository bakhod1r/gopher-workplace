# Ring Buffer

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Push` in [ringbuffer.go](ringbuffer.go):

1. If `r.size == len(r.data)`, return `errors.New("full")`.
2. `r.data[r.tail] = val`.
3. `r.tail = (r.tail + 1) % len(r.data)`.
4. `r.size++`.

## Validate

```bash
make verify
```
