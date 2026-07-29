# Parse Hex

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Fold hex digits into a value: `n = n*16 + d`, where `d` comes from `'0'..'9'`,
`'a'..'f'`, or `'A'..'F'`.

## Task

Implement `Parse(s)` (no `0x`), returning `(value, ok)`; false on non-hex/empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Parse("ff")
Output: (255, true)
```

_Explanation:_ 15*16+15

**Example 2:**

```
Input:  Parse("1A2B")
Output: (6699, true)
```

_Explanation:_ uppercase accepted

**Example 3:**

```
Input:  Parse("1g")
Output: (0, false)
```

_Explanation:_ 'g' is not a hex digit

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Hex digit value** | Map three char ranges to 0..15. |
| 2 | **Horner fold** | `n = n*16 + d`. |
| 3 | **Case folding** | Accept both `a-f` and `A-F`. |

## Hint

Per byte: if `'0'..'9'` → `c-'0'`; `'a'..'f'` → `c-'a'+10`; `'A'..'F'` →
`c-'A'+10`; else fail.

## Validate

```bash
make verify
```
