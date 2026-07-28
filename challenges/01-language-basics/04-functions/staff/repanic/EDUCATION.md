# Selective recover and re-panic

## The idea

A recover that doesn't re-raise unexpected values silently swallows genuine panics; production guards absorb only known sentinels and re-panic the rest.

## Why it matters

Blanket `recover()` is a notorious way to hide crashes and corrupt program state.

## Watch out

- Recover then inspect `r`; re-`panic(r)` what you don't recognise.
- Absorbing everything turns real bugs into silent wrong behaviour.
