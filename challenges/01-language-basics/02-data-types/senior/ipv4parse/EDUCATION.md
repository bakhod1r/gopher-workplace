# Validate before narrowing

## The idea

An IPv4 octet fits in a `byte` (0..255). Validation must happen **before** the
`byte(val)` conversion, because narrowing wraps: `byte(256) == 0`. A too-loose
bound (`> 999`) lets an invalid octet convert to a wrong byte.

## Why it matters

Config and protocol parsers convert wide accumulators into fixed-width fields
constantly. If the range check is wrong or comes after the conversion, invalid
input becomes a plausible-but-wrong value instead of an error — a real
security/robustness bug.

## Watch out

- The bound is `> 255`, exactly the byte maximum.
- Narrowing conversions never panic; they truncate silently.
- Also guard field count and empty fields, as this parser does.
