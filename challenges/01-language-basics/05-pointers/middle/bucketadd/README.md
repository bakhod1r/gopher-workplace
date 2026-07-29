# Accumulate Into Bucket Pointers

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

A map of struct POINTERS lets you mutate a bucket in place after fetching it —
unlike a map of struct values, which are not addressable.

## Task

Implement `AddScore` in [bucketadd.go](bucketadd.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AddScore(m, "a", 3); AddScore(m, "a", 4)
Output: m["a"].Total == 7
```

**Example 2:**

```
Input:  AddScore(m, "b", 5)
Output: m["b"].Total == 5
```

_Explanation:_ First touch creates the bucket.

**Example 3:**

```
Input:  AddScore(m, "a", 0)
Output: bucket exists, Total unchanged
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map of pointers** | Values are `*Bucket`, addressable via the pointer. |
| 2 | **Lazy create** | Make the bucket on first use. |
| 3 | **Mutate in place** | `m[key].Total += pts`. |

## Hint

If `m[key] == nil { m[key] = &Bucket{} }`, then `m[key].Total += pts`.

## Validate

```bash
make verify
```
