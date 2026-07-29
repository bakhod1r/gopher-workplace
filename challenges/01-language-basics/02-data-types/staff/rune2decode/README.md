# Decode 2-Byte UTF-8

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A hand-written decoder combines a 2-byte sequence into a rune. It masks the lead
byte with `0x0F` (4 bits), but a 2-byte lead carries **5** payload bits
(`110xxxxx`), so any code point ≥ U+0100 loses its top bit.

## Task

Fix the lead-byte mask between the markers in [rune2decode.go](rune2decode.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  0xC3, 0xA9
Output: 'é' (U+00E9)
```

**Example 2:**

```
Input:  0xD0, 0x81
Output: 'Ё' (U+0401)
```

_Explanation:_ Needs the 5th lead payload bit.

**Example 3:**

```
Input:  0xD8, 0xA7
Output: 'ا' (U+0627)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Payload bits** | 2-byte lead has 5 payload bits (`0x1F`). |
| 2 | **Continuation** | 6 payload bits (`0x3F`). |
| 3 | **Assembly** | `lead5<<6 |

## Hint

`rune(lead&0x1F)<<6 | rune(cont&0x3F)`.

## Validate

```bash
make verify
```
