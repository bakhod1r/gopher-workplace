# Rune Index Bounds

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A text editor addresses characters by index. The bounds check uses `len(s)` (the
**byte** length), so for multi-byte text it lets an out-of-range rune index
through — `"日本"` has 2 runes but `len` is 6.

## Task

Fix the bounds check between the markers in [runeat.go](runeat.go) to use the
rune count.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "héllo",2
Output: 'l', true
```

**Example 2:**

```
Input:  "日本",1
Output: '本', true
```

**Example 3:**

```
Input:  "日本",2
Output: 0, false
```

_Explanation:_ index out of rune range

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte vs rune length** | `len(s)` counts bytes, `len([]rune(s))` runes. |
| 2 | **Bounds check** | Compare against the rune count. |
| 3 | **Consistency** | Index and bound must use the same unit. |

## Hint

`n >= len(rs)`.

## Validate

```bash
make verify
```
