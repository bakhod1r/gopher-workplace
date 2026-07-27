# Switch Fallthrough

**Level:** junior
**Topic:** 01-language-basics → 04-conditionals
**Estimated time:** 12 min

## Context

An authorization tier is cumulative: an admin also has writer and reader
rights, a writer also has reader rights. Go switches do **not** fall through by
default, so higher tiers must opt in with an explicit `fallthrough`.

## Task

Implement `Access` in [cascade.go](cascade.go) so each level accumulates the
lower permissions via `fallthrough`: 3→"admin,write,read", 2→"write,read",
1→"read", anything else → "".

Do **not** change the function signature or the tests.

## Examples

```go
Access(3) // => "admin,write,read"
Access(2) // => "write,read"
Access(1) // => "read"
Access(0) // => ""
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **No auto fallthrough** | Unlike C, Go stops at the end of a matched case unless you say otherwise. |
| 2 | **`fallthrough`** | An explicit `fallthrough` continues into the next case's body (without re-testing it). |
| 3 | **Accumulating cases** | Placing `fallthrough` at the end of each case chains them from high to low. |

## Hint

Build up a result string. In `case 3` append `"admin,"` then `fallthrough`;
`case 2` append `"write,"` then `fallthrough`; `case 1` append `"read"`.
Missing a `fallthrough` stops the chain early.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
