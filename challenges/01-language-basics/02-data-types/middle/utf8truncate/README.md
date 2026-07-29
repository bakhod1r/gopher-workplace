# Byte-Safe Truncate

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A database column is `VARCHAR(N)` counted in **bytes**. Truncating UTF-8 text to
fit must not cut a multi-byte character in half, or you store invalid UTF-8.

## Task

Implement `Truncate(s, max)` = longest prefix ≤ `max` bytes, never splitting a
rune.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Truncate("hello", 3)
Output: "hel"
```

_Explanation:_ plain ASCII cut at 3 bytes

**Example 2:**

```
Input:  Truncate("héllo", 2)
Output: "h"
```

_Explanation:_ é is 2 bytes; including it would exceed 2 so stop at h

**Example 3:**

```
Input:  Truncate("日本", 4)
Output: "日"
```

_Explanation:_ each kanji is 3 bytes; only one fits in 4

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte vs rune width** | A rune spans 1–4 bytes. |
| 2 | **range byte index** | The index is where each rune starts. |
| 3 | **Prefix cut** | Cut at the last rune boundary ≤ max. |

## Hint

`for i := range s { if i+len(rune-at-i-bytes) > max ... }` — simpler: track the
last boundary `i` where `i <= max`; when a rune would cross `max`, stop and
return `s[:i]`.

## Validate

```bash
make verify
```
