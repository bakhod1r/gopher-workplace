# XOR vs AND-NOT

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`Revoke` uses `^` (XOR), which *toggles* bits — so revoking a permission the
user never had actually **grants** it. The intended operator is bit-clear `&^`.

## Task

Fix the single line between the markers in [perms.go](perms.go) so `Revoke`
clears bits and is a no-op for absent bits.

## Examples

```go
Revoke(Read|Write|Execute, Write) // => Read|Execute
Revoke(Read, Write)               // => Read (XOR would give Read|Write)
Revoke(all, all)                  // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **XOR toggles** | `set ^ drop` flips drop's bits, adding absent ones. |
| 2 | **AND-NOT clears** | `set &^ drop` only removes. |
| 3 | **Idempotence** | Clearing an absent bit must not change the set. |

## Hint

Change `^` to `&^`.

## Validate

```bash
make verify
```
