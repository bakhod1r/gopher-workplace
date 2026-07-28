# Shrinking Retains Pointers

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _memory-management_

## Context

Re-slicing with `s[:last]` shrinks the length but leaves the popped pointer in
the backing array, keeping the object alive. Nil the slot first: `s[last] = nil`.

## Task

Fix [shrinkretain.go](shrinkretain.go) to clear the vacated slot before shrinking.

Do **not** change the function signature or the tests.

## Examples

```go
Pop(&s) // backing slot nilled so the object can be freed
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Retention via backing array** | Shrinking doesn't clear dropped slots. |
| 2 | **Clear before shrink** | `s[last] = nil`. |
| 3 | **GC eligibility** | Nil-ing drops the last reference. |

## Hint

Clear the slot then shrink: `s[last] = nil; *sp = s[:last]`.

## Validate

```bash
make verify
```
