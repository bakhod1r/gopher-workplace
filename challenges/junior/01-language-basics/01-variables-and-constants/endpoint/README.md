# Derived Endpoint

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants
**Estimated time:** 8 min

## Context

An API client pins its host and version at build time, then builds every request
URL from them. The host and version are `const` — fixed, known to the compiler.
The assembled root is a package-level `var`, because it is *computed* from them.

The point is that `Root` must be **derived**, not retyped. Paste the finished
URL and the two constants become decoration: bump `Version` to `"v3"` and the
client keeps calling `v2`.

## Task

Implement [endpoint.go](endpoint.go) so that:

1. `BaseURL` is `"https://api.example.com"` and `Version` is `"v2"`.
2. `Root` is a package-level `var` **built from those constants** —
   `BaseURL + "/" + Version` — not a hand-written string.
3. `Path(resource)` returns `Root`, a slash, then the resource.
4. `Path("")` returns `Root` unchanged — no trailing slash.

Do **not** change the function signature or the tests.

## Examples

```go
BaseURL         // => "https://api.example.com"
Version         // => "v2"
Root            // => "https://api.example.com/v2"

Path("users")   // => "https://api.example.com/v2/users"
Path("orders")  // => "https://api.example.com/v2/orders"
Path("users/42")// => "https://api.example.com/v2/users/42"
Path("")        // => "https://api.example.com/v2"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Package-level `var`** | Lives for the whole program and is initialised before `main` runs — so it may be computed, not just literal. |
| 2 | **`const` vs `var`** | A `const` is fixed at compile time and cannot be assigned to; a `var` holds a value that code may compute or change. |
| 3 | **Constant expressions** | `BaseURL + "/" + Version` is folded by the compiler — deriving costs nothing at run time. |
| 4 | **Single source of truth** | Derive values from one definition so a change in `Version` reaches every use; a pasted copy silently goes stale. |

## Hint

`Root` is declared with `var`, not `const`, but its initialiser can still be an
expression: `var Root = BaseURL + "/" + Version`. Both operands are constants, so
the compiler folds the whole thing — you pay nothing for the readability.

For `Path`, mind the empty case: blindly appending `"/" + resource` gives a
trailing slash for `Path("")`, which the tests reject.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
