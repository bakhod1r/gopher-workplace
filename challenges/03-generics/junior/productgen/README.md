# Product

**Level:** junior  
**Topic:** 03-generics

## Context

A capacity planner multiplies scaling factors together. An empty factor list must leave the input unchanged.

## Task

Implement the stub(s) in [productgen.go](productgen.go):

1. Implement `Product`, returning the product of all elements.
2. Return `1` for an empty slice — the multiplicative identity.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Product([]int{2, 3})
Output: 6
```

**Example 2:**

```
Input:  Product([]float64{2, 0.5})
Output: 1
```

**Example 3:**

```
Input:  Product([]int{})
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Identity elements** | The empty product is `1`, not `0` — starting at zero would zero out every result. |
| 2 | **Constraints permit operations** | A type parameter can only do what every type in its set can do — that is why `+` needs a numeric constraint. |
| 3 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |

## Hint

`var out T = 1` — the literal `1` converts to any numeric `T`.

## Validate

```bash
make verify
```
