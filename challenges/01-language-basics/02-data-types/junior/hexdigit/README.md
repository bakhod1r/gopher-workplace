# Hex Digit

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Hex characters are contiguous: `'0'..'9'` then `'a'..'f'`. You reach them with
byte arithmetic on character literals.

## Task

Implement `Digit(n)` returning the lowercase hex char for `0..15`, else `'?'`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Digit(0)
Output: '0'
```

_Explanation:_ '0'+0.

**Example 2:**

```
Input:  Digit(10)
Output: 'a'
```

_Explanation:_ 'a'+(10-10).

**Example 3:**

```
Input:  Digit(15)
Output: 'f'
```

_Explanation:_ 'a'+5.

**Example 4:**

```
Input:  Digit(16)
Output: '?'
```

_Explanation:_ Out of 0..15 range.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Char arithmetic** | `'0' + n` yields the n-th digit character. |
| 2 | **byte type** | ASCII fits in a `byte` (uint8). |
| 3 | **Range guard** | Reject n outside 0..15. |

## Hint

`'0' + byte(n)` for 0–9; `'a' + byte(n-10)` for 10–15.

## Validate

```bash
make verify
```
