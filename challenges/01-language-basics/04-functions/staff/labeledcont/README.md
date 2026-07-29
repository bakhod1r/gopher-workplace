# Labeled Continue

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`break Rows` exits the OUTER loop entirely, so scanning stops at the first row
with a negative and all later rows go uncounted. The intent is to skip just the
current row and keep going: `continue Rows`.

## Task

Fix [labeledcont.go](labeledcont.go) so a negative skips only that row, not the rest.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CleanRows([[1 2 3],[1 -1 2],[4 5]])
Output: 2
```

**Example 2:**

```
Input:  rows with a negative are skipped
Output: true
```

**Example 3:**

```
Input:  count of clean rows
Output: 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Labeled continue vs break** | `continue Rows` resumes the outer loop; `break Rows` ends it. |
| 2 | **Nested-loop control** | The label targets the enclosing loop. |
| 3 | **Skip past the increment** | `continue Rows` jumps over `count++` for that row only. |

## Hint

Skip the row, don't stop: use `continue Rows` instead of `break Rows`.

## Validate

```bash
make verify
```
