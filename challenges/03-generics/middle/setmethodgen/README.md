# Set With Operations

**Level:** middle  
**Topic:** 03-generics

## Context

Permission checks combine role sets constantly, and the expressions read better as methods than as free functions.

## Task

Implement the stub(s) in [setmethodgen.go](setmethodgen.go):

1. Implement `NewSet`, `Union`, `Intersect`, and `Len`.
2. Both operations return a **new** set and leave the receivers untouched.
3. The results must chain: `a.Union(b).Intersect(c)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewSet(1,2).Union(NewSet(2,3)).Len()
Output: 3
```

**Example 2:**

```
Input:  NewSet(1,2).Intersect(NewSet(2,3)).Len()
Output: 1
```

**Example 3:**

```
Input:  NewSet[int]().Len()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods returning the same generic type** | `*Set[T]` in the result is what enables chaining. |
| 2 | **Immutability by convention** | Returning new sets keeps operands reusable across expressions. |
| 3 | **Generic types** | The type parameter belongs to the type; methods reuse it, never add to it. |

## Hint

Build a fresh set inside each operation — never write into a receiver.

## Validate

```bash
make verify
```
