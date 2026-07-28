# Typed Recover Value

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

The recovered value has type `any`; to return it as an error you must assert
`r.(error)`, not `r.(string)`. Asserting the wrong type never matches an error
panic, so `err` stays nil.

## Task

Fix the type assertion in [recovertype.go](recovertype.go).

Do **not** change the function signature or the tests.

## Examples

```go
Call(func(){ panic(errBoom) })   // => errBoom
Call(func(){ panic("x") })       // => nil
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Type assertion on recover** | `r.(error)` extracts an error panic. |
| 2 | **comma-ok assertion** | `e, ok := r.(error)` avoids a panic on mismatch. |
| 3 | **Selective conversion** | Only error panics become the return. |

## Hint

Assert the error type: `if e, ok := r.(error); ok { err = e }`.

## Validate

```bash
make verify
```
