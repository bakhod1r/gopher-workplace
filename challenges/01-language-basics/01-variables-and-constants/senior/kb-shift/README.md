# Byte-Scale Shift

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`1 << iota` gives 2,4,8 — not 1024,1048576. Binary units need the shift to jump
10 bits per step: `1 << (10 * iota)`.

## Task

Fix the single line between the markers in [units.go](units.go) so
`KB=1024, MB=1048576, GB=1073741824`.

## Examples

```go
KB // => 1024
MB // => 1048576
GB // => 1073741824
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota position** | The `_` line consumes iota==0; KB is at iota==1. |
| 2 | **Shift scaling** | `1 << (10*iota)` multiplies by 1024 per step. |
| 3 | **Implicit repeat** | MB, GB inherit the KB expression. |

## Hint

`KB ByteSize = 1 << (10 * iota)`.

## Validate

```bash
make verify
```
