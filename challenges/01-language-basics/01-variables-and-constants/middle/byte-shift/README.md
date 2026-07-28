# Binary Byte Units

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Storage sizes are powers of 1024. `iota` plus a left shift generates KB..TB in
one block — the classic Go idiom.

## Task

In [units.go](units.go):

1. Define `KB, MB, GB, TB` as `1 << (10 * iota)` (skip the `iota==0` slot with
   `_`).
2. Implement `Humanize(n)` returning the largest unit dividing `n` evenly and
   its symbol; `(n, "B")` below KB.

## Examples

```go
KB                 // => 1024
Humanize(2*KB)     // => 2, "KB"
Humanize(512)      // => 512, "B"
Humanize(5*TB)     // => 5, "TB"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota with shift** | `1 << (10*iota)` gives 1,1024,1048576,… when the first slot is skipped. |
| 2 | **Blank identifier in const** | `_ = iota` consumes value 0 so KB lands on iota==1. |
| 3 | **Untyped→typed** | `ByteSize` typing flows through the shift expression. |

## Hint

Only the first line needs the full `1 << (10 * iota)` expression; the rest
repeat it implicitly as iota climbs.

## Validate

```bash
make verify
```
