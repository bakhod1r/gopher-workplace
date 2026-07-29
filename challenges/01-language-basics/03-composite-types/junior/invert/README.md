# Invert a Map

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Building a reverse lookup: given name→id, produce id→name.

## Task

Implement `Invert(m)` returning value→key (values unique).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  {"one":1,"two":2,"three":3}
Output: {1:"one",2:"two",3:"three"}
```

**Example 2:**

```
Input:  {}
Output: {} (empty)
```

**Example 3:**

```
Input:  {"a":5}
Output: {5:"a"}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **make with new types** | Result key/value types swap. |
| 2 | **Range k,v** | Read both and store `out[v]=k`. |
| 3 | **Uniqueness** | Duplicate values would collide. |

## Hint

`out := make(map[int]string); for k, v := range m { out[v] = k }`.

## Validate

```bash
make verify
```
