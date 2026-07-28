# Convert Between Same-Layout Structs

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

Two structs with identical field layout can be reinterpreted directly:
`*(*Vec)(unsafe.Pointer(p))`. Building a Vec from only `p.X` drops `Y`.

## Task

Fix [structconv.go](structconv.go) to reinterpret the whole struct.

Do **not** change the function signature or the tests.

## Examples

```go
ToVec(&Point{3, 4}) // => {3 4}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Same layout reinterpret** | Identical fields -> direct cast. |
| 2 | **Whole struct** | Don't rebuild field by field. |
| 3 | **unsafe.Pointer bridge** | `*(*Vec)(unsafe.Pointer(p))`. |

## Hint

Reinterpret the whole struct: `return *(*Vec)(unsafe.Pointer(p))`.

## Validate

```bash
make verify
```
