# Retry Budget

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

The HTTP client should retry a failed request 3 times. Declare the retry
constant carefully: in Go, `'3'` in single quotes is a rune, not the number 3 —
an easy way to make the retry budget wildly wrong.

## Task

Declare the `MaxRetries` constant in [retries.go](retries.go) so that:

1. `MaxRetries == 3` (an integer, not a character).
2. `Budget()` returns `4` (first attempt + 3 retries).

Do **not** change the function signature or the tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MaxRetries
Output: 3
```

**Example 2:**

```
Input:  Budget()
Output: 4
```

_Explanation:_ First attempt plus 3 retries.

**Example 3:**

```
Input:  type of MaxRetries
Output: untyped int constant
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Rune literals** | `'3'` is a rune (character) constant — its value is the Unicode code point `51`, not the number 3. |
| 2 | **Integer literals** | `3` (no quotes) is an untyped integer constant with value 3. |
| 3 | **Untyped constant math** | `1 + MaxRetries` inherits the constant's kind, so a stray rune literal silently poisons the arithmetic. |

## Hint

Single quotes make a *character* constant. `'3'` is the code point for the
digit, which is `51`. Drop the quotes so it is the integer `3`.

## Validate

```bash
make verify
```
