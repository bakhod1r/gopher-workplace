# LEB128 Varint Shift

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A protobuf reader decodes varints. Each byte holds **7** value bits, but the
decoder advances the shift by 8, so multi-byte values are wildly wrong (300
decodes as 556).

## Task

Fix the shift step between the markers in [varint.go](varint.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [0xAC, 0x02]
Output: 300, 2
```

**Example 2:**

```
Input:  [0x80, 0x01]
Output: 128, 2
```

**Example 3:**

```
Input:  [0xFF, 0xFF, 0x03]
Output: 65535, 3
```

**Example 4:**

```
Input:  [0x80]
Output: 0, 0
```

_Explanation:_ Truncated mid-varint.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Base-128 groups** | 7 payload bits per byte. |
| 2 | **Continuation bit** | High bit set = more bytes follow. |
| 3 | **Little-endian groups** | Later bytes are more significant. |

## Hint

`shift += 7`.

## Validate

```bash
make verify
```
