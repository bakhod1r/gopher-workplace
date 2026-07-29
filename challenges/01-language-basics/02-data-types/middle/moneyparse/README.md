# Parse Money to Cents

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A billing service must never store money as a float. You parse the user's
`"12.34"` straight into integer cents, exactly.

## Task

Implement `Cents(s)` → integer cents; `"7"`→700, `"3.5"`→350, `"12.34"`→1234.
Reject >2 decimals or bad format.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Cents("12.34")
Output: (1234, true)
```

_Explanation:_ 12 dollars 34 cents

**Example 2:**

```
Input:  Cents("3.5")
Output: (350, true)
```

_Explanation:_ one decimal padded to 50 cents

**Example 3:**

```
Input:  Cents("7")
Output: (700, true)
```

_Explanation:_ no decimals -> whole dollars

**Example 4:**

```
Input:  Cents("1.234")
Output: (0, false)
```

_Explanation:_ more than two decimals rejected

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Integer money** | Avoid float; accumulate whole cents. |
| 2 | **Optional fraction** | 0, 1, or 2 decimal digits; pad to 2. |
| 3 | **Validation** | Reject 3+ decimals, letters, empty. |

## Hint

Parse dollars before the `.`, then 0-2 fractional digits padded to two;
`dollars*100 + frac`.

## Validate

```bash
make verify
```
