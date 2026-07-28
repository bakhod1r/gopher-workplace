# Tagless Switch

**Level:** junior
**Topic:** 01-language-basics → 04-conditionals
**Estimated time:** 8 min

## Context

Sign classification reads cleaner as a `switch` with no tag: each case is a
boolean condition, evaluated top to bottom — the same shape as an
if/else-if chain, but flatter.

## Task

Implement `Classify` in [tagless.go](tagless.go) with a tagless `switch`:
`n < 0` → "negative", `n == 0` → "zero", otherwise "positive".

Do **not** change the function signature or the tests.

## Examples

```go
Classify(-5) // => "negative"
Classify(0)  // => "zero"
Classify(7)  // => "positive"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Tagless switch** | `switch { case cond: … }` with no expression tests each case as a boolean. |
| 2 | **Top-to-bottom** | The first true case wins; order matters just like an if-ladder. |
| 3 | **default as else** | `default` covers the remaining case (here, positive). |

## Hint

`switch { case n < 0: return "negative"; case n == 0: return "zero"; default:
return "positive" }`.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
