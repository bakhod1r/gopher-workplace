# Raw String Path

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

A Windows path is full of backslashes. In an interpreted `"..."` string each `\`
must be doubled; a raw `` `...` `` literal takes the characters verbatim.

## Task

Implement `TempPath()` returning `C:\Users\temp\log.txt` using a **raw** string
literal (backticks).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TempPath()
Output: C:\Users\temp\log.txt
```

_Explanation:_ Raw literal keeps every backslash verbatim.

**Example 2:**

```
Input:  len(TempPath())
Output: 21
```

_Explanation:_ Three literal single-byte backslashes, no escaping.

**Example 3:**

```
Input:  TempPath() equals the escaped form
Output: true
```

_Explanation:_ Equal to "C:\\Users\\temp\\log.txt".

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Raw string literal** | Backtick-delimited; backslashes are literal. |
| 2 | **No escapes** | `\t`, `\n`, `\\` are not interpreted inside backticks. |
| 3 | **Single line vs multi** | Raw strings may span lines and include quotes. |

## Hint

Return `` `C:\Users\temp\log.txt` `` — backticks, no escaping.

## Validate

```bash
make verify
```
