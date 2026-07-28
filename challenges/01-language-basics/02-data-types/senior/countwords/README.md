# Robust Word Count

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

An analytics job counts words with `strings.Split(s, " ")`, which turns every
run of spaces into empty strings — so `"  a   b  "` counts as 6, not 2.

## Task

Fix the single line between the markers in [countwords.go](countwords.go) to
count only non-empty words.

## Examples

```go
Count("  a   b  ") // => 2
Count("")          // => 0
Count("   ")       // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Split vs Fields** | `strings.Fields` splits on runs of whitespace. |
| 2 | **Empty tokens** | `Split` on a single space yields empties. |
| 3 | **Whitespace kinds** | `Fields` also handles tabs/newlines. |

## Hint

`return len(strings.Fields(s))`.

## Validate

```bash
make verify
```
