# Quote a String

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

To embed a double quote inside a double-quoted string you escape it with `\"`.

## Task

Implement `Wrap(s)` returning `s` surrounded by literal `"` characters.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Wrap("hello")
Output: "hello" (with quote bytes)
```

_Explanation:_ hello surrounded by literal quote bytes.

**Example 2:**

```
Input:  Wrap("")
Output: just two quote bytes
```

_Explanation:_ Empty input yields just the two quotes.

**Example 3:**

```
Input:  Wrap("a b")
Output: "a b" (with quote bytes)
```

_Explanation:_ Spaces preserved.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Escaped quote** | `\"` puts a quote char inside a `"..."` literal. |
| 2 | **Concatenation** | Build the result with `+`. |
| 3 | **Empty string** | Wrapping "" yields just two quote chars. |

## Hint

`return "\"" + s + "\""`.

## Validate

```bash
make verify
```
