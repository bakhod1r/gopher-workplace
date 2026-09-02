# Generic Set

**Level:** junior  
**Topic:** 03-generics

## Context

A crawler tracks visited URLs. Another service tracks seen user IDs. One set type covers both.

## Task

Implement the stub(s) in [setgen.go](setgen.go):

1. Implement `NewSet`, returning a set whose map is allocated.
2. Implement `Add`, `Has`, and `Len`.
3. Adding the same value twice must not change `Len`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewSet[int](); Add(1); Has(1)
Output: true
```

**Example 2:**

```
Input:  Add(1); Add(1); Len()
Output: 1
```

**Example 3:**

```
Input:  Has(9) on empty set
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map keys need `comparable`** | A generic type storing a map must constrain its key parameter. |
| 2 | **Constructors for generic types** | A map field must be allocated, so the zero value is not usable — hence `NewSet`. |
| 3 | **Instantiating a generic type** | `Stack[int]{}` or `NewStack[int]()` fixes `T` at the use site. |

## Hint

`NewSet[int]()` needs the explicit type argument: there is no argument to infer from.

## Validate

```bash
make verify
```
