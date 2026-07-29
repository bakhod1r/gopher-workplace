# Revoke Permissions

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Removing a permission is not `-`; it is clearing bits with the AND-NOT operator
`&^`, which leaves untouched bits alone.

## Task

In [perms.go](perms.go):

1. Define `Read, Write, Execute` with `1 << iota`.
2. Implement `Revoke(set, drop)` clearing `drop`'s bits from `set`.

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

_Explanation:_ Dropping an absent bit is a no-op.

**Example 3:**

```
Input:  Revoke(Read|Write|Execute, Read)
Output: Write|Execute
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota bit flags** | `1 << iota` → 1,2,4. |
| 2 | **AND-NOT `&^`** | `set &^ drop` clears exactly drop's set bits. |
| 3 | **Idempotence** | Dropping an absent bit changes nothing. |

## Hint

`return set &^ drop` — the `&^` operator is bit-clear in one step.

## Validate

```bash
make verify
```
