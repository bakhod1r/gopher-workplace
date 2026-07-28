# Deep-Clone Map of Slices

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Copying `out[k] = v` copies the slice **header**, so the clone's slices share
backing arrays with the original. Mutating a cloned slice corrupts the source.
(`maps.Clone` is shallow for the same reason.)

## Task

Fix the assignment between the markers in
[mapclonesharedslice.go](mapclonesharedslice.go) to copy each slice.

## Examples

```go
c := Clone(m); c["a"][0]=99 // m["a"][0] stays 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shallow map copy** | Slice values are shared headers. |
| 2 | **Deep copy** | Clone each slice value. |
| 3 | **maps.Clone limit** | One level only. |

## Hint

`out[k] = append([]int{}, v...)`.

## Validate

```bash
make verify
```
