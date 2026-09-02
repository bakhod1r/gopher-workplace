# Generic Pair

**Level:** junior  
**Topic:** 03-generics

## Context

A lookup table is built from two parallel slices. Carrying the two halves together as one value removes a whole class of index bugs.

## Task

Implement the stub(s) in [pairgen.go](pairgen.go):

1. Implement `MakePair`, returning a `Pair` holding `a` and `b`.
2. Implement `Swapped`, returning a pair with the field types and values exchanged.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MakePair(1, "a")
Output: Pair[int, string]{1, "a"}
```

**Example 2:**

```
Input:  MakePair(1, "a").Swapped()
Output: Pair[string, int]{"a", 1}
```

**Example 3:**

```
Input:  MakePair(true, 2).Swapped().First
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two type parameters on a type** | `Pair[A, B]` fixes two types at instantiation. |
| 2 | **A method returning a different instantiation** | `Swapped` returns `Pair[B, A]` — same generic type, different type arguments. |
| 3 | **Methods take no new type parameters** | Go allows type parameters on the type, never extra ones on its methods. |

## Hint

`Swapped` returns `Pair[B, A]`, not `Pair[A, B]` — the order flips in the type too.

## Validate

```bash
make verify
```
