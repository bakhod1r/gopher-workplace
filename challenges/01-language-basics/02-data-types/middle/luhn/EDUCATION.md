# The Luhn algorithm

## The idea

Scanning right to left, leave the odd positions alone and double the even ones;
if a doubled digit exceeds 9, subtract 9 (same as summing its digits). The number
is valid when the grand total is divisible by 10.

## Why it matters

Luhn guards card numbers, IMEIs, and many IDs against single-digit typos and
transpositions. It is a compact, real checksum built entirely from digit
extraction and modular arithmetic.

## Watch out

- Parity is by position **from the right**, so reverse or index carefully.
- "Subtract 9" is equivalent to summing the two digits of the doubled value.
- Reject non-digit characters up front; a stray space or letter is invalid.
