# Letter Grade

**Level:** junior
**Topic:** 01-language-basics → 04-conditionals
**Estimated time:** 10 min

## Context

A gradebook turns a numeric score into a letter. The natural tool is an
`if / else if / else` ladder — but the branch order and the boundary
comparisons decide whether an exact `90` is an `A` or slips to a `B`.

## Task

Implement `Grade` in [grade.go](grade.go): 90+ `"A"`, 80+ `"B"`, 70+ `"C"`,
60+ `"D"`, below 60 `"F"`. Boundaries are inclusive.

Do **not** change the function signature or the tests.

## Examples

```go
Grade(95) // => "A"
Grade(90) // => "A"
Grade(72) // => "C"
Grade(59) // => "F"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **if / else if / else** | Branches are tested top to bottom; the first true one wins and the rest are skipped. |
| 2 | **Branch ordering** | Check the highest threshold first, or a lower branch captures high scores too early. |
| 3 | **Inclusive boundaries** | Use `>=`, not `>`, so an exact threshold like `90` lands in the right band. |

## Hint

Order the tests high to low and compare with `>=`: `if score >= 90 { "A" }
else if score >= 80 { "B" } …`, ending in an `else` for `"F"`.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
