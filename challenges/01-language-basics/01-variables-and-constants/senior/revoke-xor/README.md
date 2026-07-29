# XOR vs AND-NOT

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`Revoke` uses `^` (XOR), which *toggles* bits — so revoking a permission the
user never had actually **grants** it. The intended operator is bit-clear `&^`.

## Task

Fix the single line between the markers in [perms.go](perms.go) so `Revoke`
clears bits and is a no-op for absent bits.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Revoke(Read|Write, Write)
Output: Read
```

**Example 2:**

```
Input:  Revoke(Read, Write)
Output: Read
```

_Explanation:_ XOR would wrongly add Write.

**Example 3:**

```
Input:  Revoke(Read|Write|Execute, Read)
Output: Write|Execute
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
