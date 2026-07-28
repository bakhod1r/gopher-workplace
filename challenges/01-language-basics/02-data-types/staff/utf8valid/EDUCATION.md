# UTF-8 lead-byte classification

## The idea

UTF-8 encodes the byte count in the lead byte's high bits:

| bytes | lead pattern | mask test |
|-------|--------------|-----------|
| 1 | `0xxxxxxx` | `c < 0x80` |
| 2 | `110xxxxx` | `c&0xE0 == 0xC0` |
| 3 | `1110xxxx` | `c&0xF0 == 0xE0` |
| 4 | `11110xxx` | `c&0xF8 == 0xF0` |

Continuation bytes are `10xxxxxx` (`c&0xC0 == 0x80`).

## Why it matters

The mask must isolate **exactly** the defining prefix. `c&0xC0 == 0xC0` is too
loose — it matches every multi-byte lead, so 3- and 4-byte characters get the
wrong length and the parser desyncs. Robust UTF-8 handling underpins all text
I/O and is a security boundary (overlong/invalid sequences).

## Watch out

- Each mask pairs a bit-width with the exact prefix value.
- Check that enough continuation bytes remain and each is `10xxxxxx`.
- Full validation also rejects overlong encodings and surrogate/range violations.
