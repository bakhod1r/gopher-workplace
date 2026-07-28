# Left Pad

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A report aligns numbers by zero-padding on the **left** (`42` → `00042`). The
code appends the pad on the right instead, so columns don't line up.

## Task

Fix the return between the markers in [padleft.go](padleft.go) to pad on the
left.

## Examples

```go
Pad("42", 5, '0') // => "00042"
Pad("x", 4, '.')  // => "...x"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pad side** | Left pad prepends fill. |
| 2 | **Width in runes** | Count runes, not bytes. |
| 3 | **No truncation** | Longer-than-width returns unchanged. |

## Hint

`return pad + s`.

## Validate

```bash
make verify
```
