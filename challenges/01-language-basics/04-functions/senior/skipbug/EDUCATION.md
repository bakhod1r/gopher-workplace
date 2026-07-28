# continue as a filter guard

## The idea

`continue` skips to the next iteration; the guard must name the elements to SKIP, so a `continue` on the wanted elements silently drops them.

## Why it matters

Inverted continue-guards are a common filtering bug that passes only when the data lacks the skipped category.

## Watch out

- `continue` should guard the unwanted case (`if !keep { continue }`).
- Continuing on the values you meant to process drops them.
