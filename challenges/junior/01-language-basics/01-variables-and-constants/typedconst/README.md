# Typed Batch Limit

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants
**Estimated time:** 10 min

## Context

A batching service caps how many items go out at once, and retries each batch a
fixed number of times. Both numbers are constants — but they behave differently,
because one carries a type and the other does not.

A **typed** constant is nailed to its type: `MaxBatch` is a `byte` everywhere it
appears, so comparing it with an `int` needs an explicit conversion. An
**untyped** constant has a *kind*, not a type, and adopts whatever type its
context needs — `Retries` multiplies a `float64` with no conversion at all.

## Task

Implement [typedconst.go](typedconst.go) so that:

1. `MaxBatch` is a **typed** constant — `byte`, value `200`. Write the type.
2. `Retries` is an **untyped** constant, value `3`. Do not write a type.
3. `Fits(n)` reports whether `n` is between `0` and `MaxBatch` inclusive.
   Negative `n` never fits.
4. `Budget(base)` returns `base` multiplied by `Retries`.

Convert in the direction that cannot lose information: widen `MaxBatch` to
`int`, do not narrow `n` to `byte`. Narrowing wraps — `byte(256)` is `0`, which
would make an absurdly large batch "fit".

Do **not** change the function signatures or the tests.

## Examples

```go
Fits(0)     // => true
Fits(200)   // => true   (exactly the limit)
Fits(201)   // => false
Fits(256)   // => false  (byte(256) == 0 would wrongly say true)
Fits(-1)    // => false

Budget(1.5) // => 4.5
Budget(1)   // => 3
Budget(0)   // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Typed constants** | `const MaxBatch byte = 200` is a `byte` at every use — mixing it with an `int` is a compile error until you convert. |
| 2 | **Untyped constants** | `const Retries = 3` has a *kind* (integer), not a type; it becomes `int`, `float64` or `rune` to suit its context. |
| 3 | **Conversion direction** | `int(aByte)` always widens safely; `byte(anInt)` silently wraps modulo 256 — pick the direction that cannot lose information. |
| 4 | **Constant vs variable** | Both constants are fixed at compile time; the conversion happens where they are *used*, not where they are declared. |

## Hint

`n <= MaxBatch` will not compile: `n` is an `int`, `MaxBatch` is a `byte`. You
must convert one side — and the choice matters. Converting `n` to a `byte` makes
the comparison wrap: `byte(256)` is `0`, so a 256-item batch would "fit". Widen
the other way instead, `int(MaxBatch)`, and nothing wraps.

`Retries` needs no conversion anywhere; if you find yourself writing
`float64(Retries)`, you gave it a type it should not have.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
