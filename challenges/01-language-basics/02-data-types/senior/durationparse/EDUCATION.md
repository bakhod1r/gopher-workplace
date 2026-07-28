# Unit conversion factors

## The idea

Parse a number, then multiply by its unit's factor in the common base (seconds):
hours ×3600, minutes ×60, seconds ×1. A wrong factor silently rescales every
value that uses that unit.

## Why it matters

Timeouts, rate windows, and TTLs are parsed from strings constantly. A ×6 vs ×60
slip is invisible in code review yet makes production timeouts fire 10× early —
a real incident shape.

## Watch out

- Keep the factors exact: 60 seconds per minute, 3600 per hour.
- Require a unit after each number; a dangling number is malformed.
- Empty input is a valid zero here — decide such edge cases explicitly.
