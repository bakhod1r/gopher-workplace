# Comparison operator precision

## The idea

`>` and `>=` differ only at the boundary, but that single equal value is exactly where specifications are easy to misread.

## Why it matters

Boundary bugs slip through happy-path tests that never hit the exact threshold value.

## Watch out

- "strictly greater" ⇒ `>`; "at least" ⇒ `>=`.
- Always test the exact boundary value.
