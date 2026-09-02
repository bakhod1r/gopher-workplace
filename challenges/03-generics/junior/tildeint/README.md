# Underlying Types

**Level:** junior  
**Topic:** 03-generics

## Context

The codebase wraps raw ints in named types like `Celsius` and `Retries` for safety. A helper written for `int` alone will not accept them.

## Task

Implement the stub(s) in [tildeint.go](tildeint.go):

1. Implement `SumTemps`, returning the total of all elements.
2. The constraint `IntLike` is already declared with `~int` — study why the `~` is required.
3. Return the zero value for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumTemps([]Celsius{1, 2})
Output: Celsius(3)
```

**Example 2:**

```
Input:  SumTemps([]int{5})
Output: 5
```

**Example 3:**

```
Input:  SumTemps([]Celsius{})
Output: Celsius(0)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `~` token** | `~int` means "any type whose underlying type is int", so named types like `type Celsius int` are included. |
| 2 | **Declaring a constraint** | A constraint is an interface holding a type set instead of methods. |
| 3 | **Named types** | Reused from language basics: `type Celsius int` is a distinct type with `int` as its underlying type. |

## Hint

The body is a plain sum; the lesson lives in the `~` inside the constraint.

## Validate

```bash
make verify
```
