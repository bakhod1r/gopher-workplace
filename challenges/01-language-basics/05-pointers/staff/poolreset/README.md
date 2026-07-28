# Reset Forgets a Field

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _memory-management_

## Context

Resetting for reuse must clear ALL state. Zeroing only Len leaves Data holding
stale bytes; also truncate the slice with `b.Data = b.Data[:0]` (keeping capacity
for reuse).

## Task

Fix [poolreset.go](poolreset.go) to fully reset the buffer.

Do **not** change the function signature or the tests.

## Examples

```go
Reset(&Buf{Data: [5]byte, Len: 5}) // Len 0, Data len 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Reset all fields** | Both Len and Data. |
| 2 | **Truncate keeping capacity** | `b.Data = b.Data[:0]`. |
| 3 | **Reuse safety** | Stale data must not leak into the next use. |

## Hint

Clear both: `b.Len = 0; b.Data = b.Data[:0]`.

## Validate

```bash
make verify
```
