# Triple

**Level:** junior  
**Topic:** 03-generics

## Context

A join step carries a key, a left row, and a right row through a pipeline as one value.

## Task

Implement the stub(s) in [triplegen.go](triplegen.go):

1. Implement `MakeTriple`, returning a triple of the three arguments.
2. Implement `Rotated`, returning `(b, c, a)` with the type arguments rotated to match.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MakeTriple(1, "a", true)
Output: Triple[int, string, bool]
```

**Example 2:**

```
Input:  MakeTriple(1, "a", true).Rotated().First
Output: "a"
```

**Example 3:**

```
Input:  MakeTriple(1, "a", true).Rotated().Third
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Three type parameters** | Parameters are positional: rotating the values must rotate the types too. |
| 2 | **Methods returning re-parameterised types** | `Triple[B, C, A]` is a different instantiation of the same type. |
| 3 | **Methods take no new type parameters** | Go allows type parameters on the type, never extra ones on its methods. |

## Hint

Rotate the type arguments in the return type exactly as you rotate the values.

## Validate

```bash
make verify
```
