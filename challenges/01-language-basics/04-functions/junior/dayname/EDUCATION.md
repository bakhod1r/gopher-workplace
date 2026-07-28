# Value switch for lookups

## The idea

`switch x { case a: }` compares `x` against each case value and runs the first equal one; `default` handles the rest.

## Why it matters

Small fixed mappings read more clearly as a switch than as a map or if-chain.

## Watch out

- Provide a `default` so unexpected input has a defined result.
- Cases can group values: `case 0, 6:` for weekend logic.
