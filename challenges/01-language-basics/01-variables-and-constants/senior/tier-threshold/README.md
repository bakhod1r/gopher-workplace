# Tier Threshold Off-By-One

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`iota * 100` makes Bronze `0` — a free tier nobody has to earn. The thresholds
must start at 100.

## Task

Fix the single line between the markers in [tiers.go](tiers.go) so
`Bronze=100, Silver=200, Gold=300`.

## Examples

```go
Bronze // => 100
Silver // => 200
Gold   // => 300
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota starts at 0** | `iota*100` → 0,100,200. |
| 2 | **Offset expression** | `(iota+1)*100` → 100,200,300. |
| 3 | **Implicit repeat** | Silver, Gold inherit the Bronze expression. |

## Hint

`Bronze Tier = (iota + 1) * 100`.

## Validate

```bash
make verify
```
