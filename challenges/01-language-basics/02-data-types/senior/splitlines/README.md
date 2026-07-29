# CRLF-Safe Line Split

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A log parser splits on `\n` but forgets that Windows files use `\r\n`, so every
line keeps a trailing `\r`. Comparisons and trims downstream then mysteriously
fail.

## Task

Fix the loop body between the markers in [splitlines.go](splitlines.go) to strip
a trailing `\r`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "a\nb\nc"
Output: ["a" "b" "c"]
```

**Example 2:**

```
Input:  "a\r\nb\r\nc"
Output: ["a" "b" "c"]
```

_Explanation:_ CRLF trimmed

**Example 3:**

```
Input:  ""
Output: [""]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Line endings** | LF vs CRLF across platforms. |
| 2 | **TrimSuffix** | `strings.TrimSuffix(p, "\r")`. |
| 3 | **Idempotence** | LF input already has no `\r` to strip. |

## Hint

`parts[i] = strings.TrimSuffix(p, "\r")`.

## Validate

```bash
make verify
```
