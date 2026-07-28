# Struct Field Padding

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Fields are laid out in declaration order with alignment padding. `bool, int64,
bool` forces the `int64` to an 8-byte boundary and pads the trailing bool, giving
24 bytes. Grouping the wide field first packs it to 16.

## Task

Reorder the fields between the markers in
[structpadding.go](structpadding.go) to minimize size (16 bytes on 64-bit).

## Examples

```go
unsafe.Sizeof(Record{}) // must be 16
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Alignment** | An int64 aligns to 8 bytes. |
| 2 | **Padding** | Gaps inserted to satisfy alignment. |
| 3 | **Field ordering** | Largest-first reduces padding. |

## Hint

Order `B int64` first, then the two bools: `B; A; C`.

## Validate

```bash
make verify
```
