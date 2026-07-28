# Min Pointer Bad Seed

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Seeding min with the LAST element breaks the earliest-tie rule: when the minimum
appears earlier, the strict `<` comparison won't replace an equal last element.
Seed from the first element instead.

## Task

Fix the seed in [minseed.go](minseed.go) so ties return the earliest minimum.

Do **not** change the function signature or the tests.

## Examples

```go
MinPtr([3 1 4 1 5]) // => &xs[1]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Seed from the front** | `min := &xs[0]`. |
| 2 | **Strict comparison + tie rule** | `<` keeps the earliest on ties. |
| 3 | **Consistent scan** | Seed and loop must agree. |

## Hint

Seed from the first element: `min := &xs[0]`.

## Validate

```bash
make verify
```
