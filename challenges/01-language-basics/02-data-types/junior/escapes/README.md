# Escape Sequences

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

TSV output uses a literal tab between fields and a newline at the end. Both are
escape sequences inside an interpreted string.

## Task

Implement `Row(name, value)` returning `name`, a tab, `value`, and a newline.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Row("id", "42")
Output: "id\t42\n"
```

_Explanation:_ name, a tab byte, value, then a newline byte.

**Example 2:**

```
Input:  Row("a", "b")
Output: "a\tb\n"
```

_Explanation:_ Same TSV shape.

**Example 3:**

```
Input:  Row("", "")
Output: "\t\n"
```

_Explanation:_ Empty fields still get the tab and newline.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Interpreted literal** | Double quotes; escapes are processed. |
| 2 | **`\t` and `\n`** | Tab and newline as single characters. |
| 3 | **Concatenation** | `+` joins strings. |

## Hint

`return name + "\t" + value + "\n"`.

## Validate

```bash
make verify
```
