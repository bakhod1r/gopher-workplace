# If With Init

**Level:** junior
**Topic:** 01-language-basics → 04-conditionals
**Estimated time:** 10 min

## Context

A sharding helper buckets values by `n % 3`. Go lets you compute that remainder
right in the `if` header with an init statement, keeping the variable scoped to
the if/else chain instead of leaking into the surrounding function.

## Task

Implement `Bucket` in [ifinit.go](ifinit.go) using an `if` with an init clause
(`if r := n % 3; r == 0 { … }`): remainder 0 → "zero", 1 → "one", 2 → "two".

Do **not** change the function signature or the tests.

## Examples

```go
Bucket(9)  // => "zero"
Bucket(10) // => "one"
Bucket(11) // => "two"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **if with init** | `if x := expr; cond { … }` runs the init, then tests the condition. |
| 2 | **Init scope** | The init variable is visible only inside the if/else-if/else chain, not after it. |
| 3 | **Modulo** | `n % 3` yields 0, 1, or 2 for non-negative `n`. |

## Hint

`if r := n % 3; r == 0 { return "zero" } else if r == 1 { return "one" } else
{ return "two" }`. `r` lives only inside the chain.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
