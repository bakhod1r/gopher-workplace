# Hex Encode

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Each byte becomes two hex chars: the high nibble `b>>4` and the low nibble
`b&0x0f`, each mapped through `"0123456789abcdef"`.

## Task

Implement `Encode(b)` (lowercase, two chars/byte). No `encoding/hex`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Encode([]byte{0xFF})
Output: "ff"
```

_Explanation:_ high nibble f, low nibble f

**Example 2:**

```
Input:  Encode([]byte{0x1a,0x2b})
Output: "1a2b"
```

_Explanation:_ two bytes -> four hex chars

**Example 3:**

```
Input:  Encode([]byte("Go"))
Output: "476f"
```

_Explanation:_ 'G'=0x47, 'o'=0x6f

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nibble split** | `b>>4` and `b&0x0f`. |
| 2 | **Digit table** | Index `"0..f"` by nibble value. |
| 3 | **Builder** | Append two chars per byte. |

## Hint

`const hexd = "0123456789abcdef"`; append `hexd[b>>4]` and `hexd[b&0x0f]`.

## Validate

```bash
make verify
```
