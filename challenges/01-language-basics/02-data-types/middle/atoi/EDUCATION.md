# Parsing integers by hand

## The idea

Each digit shifts the running value up one decimal place: `n = n*10 + (c-'0')`.
`c - '0'` converts a digit character to its numeric value because ASCII digits
are contiguous.

## Why it matters

It is the core of every `Atoi`/`ParseInt`, and it shows the interplay of bytes,
character arithmetic, and accumulation. Explicit validation teaches the
`(value, ok)` error style before `error` is introduced.

## Watch out

- Range over **bytes** here (ASCII digits); ranging runes also works but bytes
  are simpler for digits.
- Handle the sign once, after (or before) the digit loop — not per digit.
- Real parsers also detect overflow; this simplified one does not.
