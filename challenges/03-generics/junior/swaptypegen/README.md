# Swap Fields In Place

**Level:** junior  
**Topic:** 03-generics

## Context

A range widget stores a start and an end that the user can drag past each other. The stored pair must be normalisable on demand.

## Task

Implement the stub(s) in [swaptypegen.go](swaptypegen.go):

1. Implement `Swap`, exchanging the two stored values in place.
2. Implement `Ordered`, returning the two values smallest first without modifying the receiver.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SamePair[int]{1, 2}.Swap()
Output: pair becomes {2, 1}
```

**Example 2:**

```
Input:  SamePair[int]{2, 1}.Ordered()
Output: 1, 2
```

**Example 3:**

```
Input:  SamePair[int]{1, 2}.Ordered()
Output: 1, 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer versus value receivers** | `Swap` mutates so it needs `*SamePair[T]`; `Ordered` only reads. |
| 2 | **One parameter, two fields** | `SamePair[T]` forces both values to share a type — unlike `Pair[A, B]`. |
| 3 | **Constrained type parameter** | `cmp.Ordered` is what makes `Ordered` possible at all. |

## Hint

Mutating method takes a pointer receiver; reading method takes a value receiver.

## Validate

```bash
make verify
```
