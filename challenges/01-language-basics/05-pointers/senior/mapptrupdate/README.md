# Update Map Struct via Pointer

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

`acc := *a` copies the account; mutating `acc` changes the copy, not the stored
object. Mutate through the pointer: `a.Balance += amt`.

## Task

Fix [mapptrupdate.go](mapptrupdate.go) so the stored account is credited.

Do **not** change the function signature or the tests.

## Examples

```go
Credit(m, 1, 50) // m[1].Balance += 50
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Copy vs pointer mutation** | `*a` copies the struct. |
| 2 | **Mutate through the pointer** | `a.Balance += amt`. |
| 3 | **Stored object** | The map holds the pointer, so edits stick. |

## Hint

Mutate the pointee directly: `a.Balance += amt`.

## Validate

```bash
make verify
```
