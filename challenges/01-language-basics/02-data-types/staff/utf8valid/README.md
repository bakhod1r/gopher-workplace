# UTF-8 Lead Byte Mask

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A protocol layer validates incoming UTF-8. The 2-byte lead test uses
`c&0xC0 == 0xC0`, which also matches 3- and 4-byte leads (0xE0, 0xF0), so those
are decoded as 2-byte sequences and valid text is rejected.

## Task

Fix the 2-byte lead mask between the markers in
[utf8valid.go](utf8valid.go) to match only `110xxxxx`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "é" (C3 A9)
Output: true
```

**Example 2:**

```
Input:  "€" (E2 82 AC)
Output: true
```

**Example 3:**

```
Input:  [0xC3]
Output: false
```

_Explanation:_ Truncated 2-byte sequence.

**Example 4:**

```
Input:  [0x80]
Output: false
```

_Explanation:_ Lone continuation byte.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **UTF-8 lead bytes** | 110x=2, 1110=3, 11110=4 bytes. |
| 2 | **Prefix masks** | 2-byte lead = `c&0xE0 == 0xC0`. |
| 3 | **Continuation** | Each following byte is `10xxxxxx`. |

## Hint

`case c&0xE0 == 0xC0:` (top three bits are `110`).

## Validate

```bash
make verify
```
