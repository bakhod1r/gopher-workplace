# Max By Key

**Level:** junior  
**Topic:** 03-generics

## Context

A roster picks the highest scorer. The elements themselves are structs, which have no ordering of their own.

## Task

Implement the stub(s) in [maxbykey.go](maxbykey.go):

1. Implement `MaxBy`, returning the element whose `key(v)` is largest and `true`.
2. On a tie keep the earlier element.
3. Return the zero value and `false` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MaxBy([]P{{"a", 1}, {"b", 3}}, score)
Output: P{"b", 3}, true
```

**Example 2:**

```
Input:  MaxBy([]P{{"a", 2}, {"b", 2}}, score)
Output: P{"a", 2}, true
```

**Example 3:**

```
Input:  MaxBy([]P{}, score)
Output: P{}, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Separating element and key types** | `T any` because elements are never compared; `K cmp.Ordered` because keys are. |
| 2 | **Key extraction functions** | A `func(T) K` decouples "what to compare" from "how to scan". |
| 3 | **Caching the key** | Storing `bestKey` avoids calling `key` again on every iteration. |

## Hint

Two type parameters: one for the element, one for the ordered key.

## Validate

```bash
make verify
```
