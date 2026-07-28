# Map Comma-Ok

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types
**Estimated time:** 10 min

## Context

A feature-flag store maps names to counts. A plain `m[key]` returns `0` both
when a key is missing *and* when it is present with value `0` — the code can't
tell them apart. Go's comma-ok read solves this.

## Task

Implement `Lookup` in [lookup.go](lookup.go) so it returns the value for `key`
and a boolean reporting whether the key was actually present. A key mapped to
`0` must report `true`.

Do **not** change the function signature or the tests.

## Examples

```go
Lookup(map[string]int{"a": 5}, "a") // => 5, true
Lookup(map[string]int{"a": 5}, "z") // => 0, false
Lookup(map[string]int{"z": 0}, "z") // => 0, true   (present, value 0)
Lookup(nil, "a")                    // => 0, false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok read** | `v, ok := m[key]` returns the value and whether the key exists. |
| 2 | **Zero value ambiguity** | Plain `m[key]` yields the value type's zero for a missing key — indistinguishable from a stored zero. |
| 3 | **Nil map reads** | Reading a `nil` map is safe and always reports "not present". |

## Hint

Use the two-result form of the map index: `v, ok := m[key]` and return both.
Don't test `v != 0` to decide presence — that misreports a stored `0`.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
