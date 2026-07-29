# Flatten Slices

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

A variadic of slices (`...[]int`) accepts any number of groups; flattening
appends each group's elements into one result.

## Task

Implement `Flatten` in [flatten.go](flatten.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Flatten([1 2], [3], [4 5])
Output: [1 2 3 4 5]
```

**Example 2:**

```
Input:  Flatten()
Output: []
```

**Example 3:**

```
Input:  Flatten(nil, [9])
Output: [9]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic of slices** | `groups ...[]int`. |
| 2 | **Spread each group** | `append(out, g...)`. |
| 3 | **Order preserved** | Groups concatenate left to right. |

## Hint

`for _, g := range groups { out = append(out, g...) }; return out`.

## Validate

```bash
make verify
```
