# Base64 Digit Offset

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A base64 decoder maps characters to 6-bit values. Digits `0-9` are values
**52-61**, but the code adds 53, so every digit is off by one and decoded bytes
are corrupt.

## Task

Fix the digit offset between the markers in [base64val.go](base64val.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  'A'
Output: 0, true
```

**Example 2:**

```
Input:  'z'
Output: 51, true
```

**Example 3:**

```
Input:  '0'
Output: 52, true
```

_Explanation:_ Digits start at value 52.

**Example 4:**

```
Input:  '='
Output: 0, false
```

_Explanation:_ Not a base64 char.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Base64 alphabet** | A-Z=0-25, a-z=26-51, 0-9=52-61, +=62, /=63. |
| 2 | **Range offsets** | Each run starts where the previous ended. |
| 3 | **Off-by-one** | 'a' is 26 (after 25 letters), '0' is 52. |

## Hint

`return int(c-'0') + 52`.

## Validate

```bash
make verify
```
