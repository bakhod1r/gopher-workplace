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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m={a:[1 2], b:[3]}, then c[a][0]=99
Output: m[a][0] stays 1
```

**Example 2:**

```
Input:  len(clone)
Output: 2
```

**Example 3:**

```
Input:  clone[b][0]
Output: 3
```

_Explanation:_ values are independent copies.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
