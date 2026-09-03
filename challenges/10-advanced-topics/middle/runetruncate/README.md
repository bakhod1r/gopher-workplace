# Cut A String Without Splitting A Character

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A log line is cut to 200 bytes for display. Non-ASCII messages end in a broken character, and the terminal shows a replacement glyph for the last one.

## Task

Implement [runetruncate.go](runetruncate.go):

1. Return the longest prefix of `s` that is at most `n` bytes and ends on a character boundary.
2. `n <= 0` returns the empty string; `n >= len(s)` returns `s`.
3. The result must be a substring — zero allocations.

Replace the stub body in [runetruncate.go](runetruncate.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Truncate("hello", 3)
Output: "hel"
```

**Example 2:**

```
Input:  Truncate("héllo", 2)
Output: "h"
```

_Explanation:_ Cutting at 2 would split é, so back up.

**Example 3:**

```
Input:  Truncate("hi", 9)
Output: "hi"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **UTF-8 continuation bytes** | Every byte of a character after the first has the top bits 10. |
| 2 | **utf8.RuneStart** | Reports whether a byte can begin a character. |
| 3 | **Substrings do not copy** | Backing up an index costs nothing. |

## Hint

Walk `n` backwards while `s[n]` is not the start of a character.

## Validate

```bash
make verify
```
