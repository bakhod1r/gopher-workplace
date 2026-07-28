# Composing layered rules

## The idea

Nested exceptions collapse into a single boolean expression when you order the divisibility tests by precedence.

## Why it matters

Calendar, tax, and eligibility rules are full of exception-on-exception logic like this.

## Watch out

- 1900 is divisible by 4 and 100 but not 400, so it is NOT a leap year.
- Parenthesise to keep `&&` binding tighter than `||`.
