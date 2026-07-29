# Base Conversion

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Writing a number in base b is repeated division: the remainders are the digits,
least significant first, so you reverse at the end.

## Task

Implement `Format(n, base)` for base 2..16, lowercase digits. `Format(0,b)="0"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Format(255, 16)
Output: "ff"
```

_Explanation:_ 255 = 15*16+15 -> digits f,f

**Example 2:**

```
Input:  Format(5, 2)
Output: "101"
```

_Explanation:_ 5 in binary

**Example 3:**

```
Input:  Format(0, 2)
Output: "0"
```

_Explanation:_ zero is the special case

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Repeated division** | `n%base` is the next digit, `n/=base`. |
| 2 | **Digit mapping** | 0-9 then a-f via char arithmetic. |
| 3 | **Reverse order** | Digits come out least-significant first. |

## Hint

Collect `"0123456789abcdef"[n%base]`, divide, then reverse the bytes.

## Validate

```bash
make verify
```
