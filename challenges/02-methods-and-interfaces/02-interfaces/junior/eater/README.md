# Eater

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A zoo feeding schedule asks each animal what it eats and whether a given food is acceptable.

## Task

Implement the stub(s) in [eater.go](eater.go):

1. Implement `Eats` on `Cow` — only `"grass"`.
2. Implement `Eats` on `Lion` — only `"meat"`.
3. Implement `FeedableCount`, which counts how many eaters accept the food.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Cow{}.Eats("grass")
Output: true
```

**Example 2:**

```
Input:  Lion{}.Eats("grass")
Output: false
```

**Example 3:**

```
Input:  FeedableCount([]Eater{Cow{}, Lion{}, Cow{}}, "grass")
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface parameters** | The method takes an argument; the contract covers signature, not just name. |
| 2 | **Polymorphic dispatch** | The same call runs different code per dynamic type. |
| 3 | **Counting loop** | Reused: conditional increment over a slice. |

## Hint

Each `Eats` is a one-line string comparison.

## Validate

```bash
make verify
```
