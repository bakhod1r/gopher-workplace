# Any and All

**Level:** junior  
**Topic:** 03-generics

## Context

A validation step asks two questions of the same batch: did anything fail, and did everything pass?

## Task

Implement the stub(s) in [anyofgen.go](anyofgen.go):

1. Implement `Any`, reporting whether at least one element satisfies `pred` (false for an empty slice).
2. Implement `All`, reporting whether every element satisfies `pred` (true for an empty slice).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Any([]int{1, 2}, isEven)
Output: true
```

**Example 2:**

```
Input:  All([]int{2, 4}, isEven)
Output: true
```

**Example 3:**

```
Input:  All([]int{}, isEven)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Vacuous truth** | `All` on an empty slice is `true` — there is no counterexample. |
| 2 | **Functions as values** | Reused from language basics: a `func(T) U` parameter is an ordinary value. |
| 3 | **Short-circuit scans** | Reused from language basics: return as soon as the answer is decided. |

## Hint

The two functions are mirror images: `Any` returns early on success, `All` returns early on failure.

## Validate

```bash
make verify
```
