# FNV-1a Hash Order

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A hash map uses FNV-1a for keys. The loop multiplies **then** XORs — that is
FNV-1, a different hash — so results never match the reference vectors and
cross-service hashing disagrees.

## Task

Fix the order between the markers in [fnv1a.go](fnv1a.go): FNV-1a is XOR first,
then multiply.

## Examples

```go
Hash([]byte(""))     // => 0x811c9dc5
Hash([]byte("a"))    // => 0xe40c292c
Hash([]byte("foobar")) // => 0xbf9cf968
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **FNV-1a step** | `h = (h ^ b) * prime`. |
| 2 | **FNV-1 vs 1a** | 1a XORs before multiplying; 1 after. |
| 3 | **Unsigned wrap** | Multiply overflows mod 2³² by design. |

## Hint

`h = (h ^ uint32(b)) * prime32`.

## Validate

```bash
make verify
```
