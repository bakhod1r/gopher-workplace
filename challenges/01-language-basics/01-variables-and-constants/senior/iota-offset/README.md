# Off-By-One iota Flags

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A permission block compiles, `Has` looks right, yet every flag is double what it
should be — masks silently overlap. The bug is one token in the iota expression.

## Task

Fix the single line between the `CHANGE CODE` markers in [perms.go](perms.go) so
`Read=1, Write=2, Execute=4`. Do not touch anything else.

## Examples

```go
Read     // must be 1
Write    // must be 2
Execute  // must be 4
Has(Read|Write, Read) // => true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota base** | The first line's expression seeds the whole block. |
| 2 | **`1 << iota`** | Correct shift starts at `1 << 0 == 1`. |
| 3 | **Bit overlap** | Off-by-one shifts make masks collide, breaking `Has`. |

## Hint

`1 << (iota + 1)` starts Read at 2. Drop the `+ 1`.

## Validate

```bash
make verify
```
