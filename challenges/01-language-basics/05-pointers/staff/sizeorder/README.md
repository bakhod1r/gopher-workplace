# Field Order and Struct Size

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

Each field aligns to its size; the compiler inserts padding. `bool,int64,bool`
pads to 24; ordering widest-first `int64,bool,bool` packs to 16.

## Task

Fix the field order in [sizeorder.go](sizeorder.go) so `Record` is 16 bytes.

Do **not** change the function signature or the tests.

## Examples

```go
unsafe.Sizeof(Record{}) // must be 16
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Alignment and padding** | Fields align to their size. |
| 2 | **Widest-first ordering** | Put int64 before the bools. |
| 3 | **unsafe.Sizeof** | Reveals the true size. |

## Hint

Order widest-to-narrowest: `B int64; A bool; C bool`.

## Validate

```bash
make verify
```
