# Struct Copy Shares Slice

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

Returning `d` copies the struct — but the `Tags` **slice header** is copied, so
both share the same backing array. Mutating the clone's tags corrupts the
original.

## Task

Fix `Clone` between the markers in [shallowclone.go](shallowclone.go) to copy the
`Tags` slice independently.

## Examples

```go
c := Clone(orig); c.Tags[0]="MUT" // orig.Tags unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shallow struct copy** | Value copy duplicates fields, not referents. |
| 2 | **Slice field aliasing** | The copied header shares the array. |
| 3 | **Deep copy** | Clone the slice explicitly. |

## Hint

Copy the struct, then `d.Tags = append([]string{}, d.Tags...)` (or slices.Clone).

## Validate

```bash
make verify
```
