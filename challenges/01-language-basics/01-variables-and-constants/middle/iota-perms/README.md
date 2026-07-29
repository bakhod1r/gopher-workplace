# Permission Bits

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A filesystem stores permissions as bit flags packed into one byte. You define
the flags with `iota` and test membership with bitwise AND.

## Task

In [perms.go](perms.go):

1. Define `Read`, `Write`, `Execute` as consecutive power-of-two bits using
   `iota` (`Read=1`, `Write=2`, `Execute=4`).
2. Implement `Has(set, want)` returning true only if `set` contains every bit
   in `want`.

Do not change signatures or tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Read, Write, Execute
Output: 1, 2, 4
```

**Example 2:**

```
Input:  Has(Read|Write, Read)
Output: true
```

**Example 3:**

```
Input:  Has(Read, Write)
Output: false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota bit flags** | `1 << iota` in a const block yields 1,2,4,8… |
| 2 | **Typed constants** | Giving the first const a type applies it to the whole run. |
| 3 | **Bitwise AND mask** | `set&want == want` tests that all wanted bits are set. |

## Hint

`Read Permission = 1 << iota` then `Write` / `Execute` on their own lines reuse
the expression with iota incrementing.

## Validate

```bash
make verify
```
