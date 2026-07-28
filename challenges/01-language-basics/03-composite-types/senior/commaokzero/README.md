# Zero vs Missing

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`GetOr` checks `v != 0` to decide presence, so a key stored with value `0`
wrongly returns the default. Only the comma-ok form distinguishes them.

## Task

Fix the presence check between the markers in
[commaokzero.go](commaokzero.go) to use comma-ok.

## Examples

```go
GetOr({zero:0}, "zero", 99) // => 0 (present)
GetOr({}, "x", 99)          // => 99
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok** | `v, ok := m[key]`. |
| 2 | **Zero ≠ absent** | A stored 0 is present. |
| 3 | **Presence flag** | Branch on `ok`, not the value. |

## Hint

`if v, ok := m[key]; ok { return v }`.

## Validate

```bash
make verify
```
