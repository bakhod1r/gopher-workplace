# JSON String Escaping

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A hand-rolled JSON writer escapes quotes and control chars but forgets the
**backslash** — so `a\b` is emitted as `a\b`, producing invalid JSON that a
strict parser rejects.

## Task

Add the missing `'\\'` case between the markers in
[jsonescape.go](jsonescape.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a"b
Output: a\"b
```

**Example 2:**

```
Input:  a\b
Output: a\\b
```

_Explanation:_ backslash must be escaped

**Example 3:**

```
Input:  line\n
Output: line\n
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Escape set** | `\`, `"`, `\n`, `\t`, `\r` at minimum. |
| 2 | **Backslash first** | It is itself an escape character. |
| 3 | **Raw literals** | `` `\\` `` is a two-char string: backslash, backslash. |

## Hint

Add `case '\\': b.WriteString(`\\\\`)` (a backslash escaped as two).

## Validate

```bash
make verify
```
