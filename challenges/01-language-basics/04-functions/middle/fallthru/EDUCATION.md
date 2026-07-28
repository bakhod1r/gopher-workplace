# The fallthrough keyword

## The idea

Unlike C, Go breaks after each case; `fallthrough` opts into entering the next case body (without checking its condition).

## Why it matters

It compactly expresses cumulative/tiered logic but is easy to misuse because it skips the next case's test.

## Watch out

- `fallthrough` enters the NEXT case unconditionally — order your cases so that's correct.
- It must be the last statement in a case.
