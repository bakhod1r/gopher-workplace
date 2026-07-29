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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Doc{Title:A, Tags:[x,y]}
Output: independent copy; mutating copy.Tags[0] doesn't affect original
```

**Example 2:**

```
Input:  Doc{Title:B, Tags:[]}
Output: Doc{Title:B, Tags:[]}
```

**Example 3:**

```
Input:  Doc{Title:C, Tags:[t]}
Output: Doc{Title:C, Tags:[t]}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
