# Hex Nibble Order

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A wire-protocol decoder turns hex into bytes, but every byte comes out with its
two nibbles swapped: `"1a"` decodes to `0xa1`. The high nibble is the *first*
character.

## Task

Fix the single line between the markers in [hexdecode.go](hexdecode.go) so the
first hex char is the high nibble.

## Examples

```go
Decode("1a2b") // => [0x1a, 0x2b]
Decode("ff")   // => [0xff]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nibble position** | First char = high 4 bits. |
| 2 | **Shift + OR** | `hi<<4 | lo`. |
| 3 | **Even length** | Two hex chars per byte. |

## Hint

`byte(hi<<4 | lo)`.

## Validate

```bash
make verify
```
