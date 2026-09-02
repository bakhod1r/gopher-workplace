# Constraint With A Method

**Level:** junior  
**Topic:** 03-generics

## Context

A checkout totals line items. Every item type knows its own price in cents, and the totalling code should stay one function.

## Task

Implement the stub(s) in [constraintifcmethod.go](constraintifcmethod.go):

1. Implement `TotalCents`, summing `Cents()` across the items.
2. The `Priced` constraint is provided: it requires a `Cents() int` method.
3. Return `0` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TotalCents([]Book{{200}, {350}})
Output: 550
```

**Example 2:**

```
Input:  TotalCents([]Book{})
Output: 0
```

**Example 3:**

```
Input:  TotalCents([]Coffee{{450}})
Output: 450
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method constraints** | A constraint interface may list methods, type sets, or both. |
| 2 | **Static dispatch** | The call `it.Cents()` resolves per instantiation — no interface value is built. |
| 3 | **Constraint versus interface** | `Priced` could also be used as a plain interface; as a constraint it keeps the slice typed. |

## Hint

`Priced` is an ordinary interface used in the constraint position.

## Validate

```bash
make verify
```
