# Word Wrap Width

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A terminal UI wraps text to a column width, but an off-by-one (`<` instead of
`<=`) breaks one character too early, so `"aa bb"` (exactly 5) is split at
width 5.

## Task

Fix the fit test between the markers in [wordwrap.go](wordwrap.go) so a line may
use the full width.

## Examples

```go
Wrap("aa bb cc", 5)        // => ["aa bb", "cc"]
Wrap("the quick brown fox",10) // => ["the quick","brown fox"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Fit condition** | New length `<= width` fits. |
| 2 | **Account for space** | `+1` for the joining space. |
| 3 | **Off-by-one** | `<` rejects an exactly-full line. |

## Hint

`if len(line)+1+len(w) <= width`.

## Validate

```bash
make verify
```
