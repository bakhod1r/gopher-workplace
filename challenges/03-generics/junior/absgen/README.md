# Absolute Value

**Level:** junior  
**Topic:** 03-generics

## Context

A diff report shows how far a measurement drifted, without caring in which direction.

## Task

Implement the stub(s) in [absgen.go](absgen.go):

1. Implement `Abs`, returning `v` without its sign.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Abs(-3)
Output: 3
```

**Example 2:**

```
Input:  Abs(3)
Output: 3
```

**Example 3:**

```
Input:  Abs(-2.5)
Output: 2.5
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Signed constraints** | Unary `-` and comparison to `0` require signed numeric types — unsigned types must stay out of the set. |
| 2 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |
| 3 | **Declaring a constraint** | A constraint is an interface holding a type set instead of methods. |

## Hint

`-v` and `v < 0` both need a signed numeric constraint.

## Validate

```bash
make verify
```
