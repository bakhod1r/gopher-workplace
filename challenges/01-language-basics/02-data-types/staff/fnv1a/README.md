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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ""
Output: 0x811c9dc5
```

_Explanation:_ The FNV offset basis.

**Example 2:**

```
Input:  "a"
Output: 0xe40c292c
```

**Example 3:**

```
Input:  "foobar"
Output: 0xbf9cf968
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
