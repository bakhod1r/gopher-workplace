# How Big Is This, Really

**Level:** junior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A capacity estimate multiplies the record count by "about the size of the struct". The estimate is off by a factor of three because nobody measured the string and slice headers.

## Task

Implement [sizeof.go](sizeof.go):

1. Return `unsafe.Sizeof` for the `Header` type, its `Id` field and its `Name` field.
2. Do not hard-code numbers — the answers must follow the type.

Replace the stub body in [sizeof.go](sizeof.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Sizes()
Output: 40, 8, 16
```

_Explanation:_ On a 64-bit build; a string header is two words.

**Example 2:**

```
Input:  Sizeof of a long Name
Output: still 16
```

_Explanation:_ The header size does not depend on the text.

**Example 3:**

```
Input:  header vs id
Output: header is larger
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Sizeof** | A compile-time constant giving the type's size in bytes. |
| 2 | **Headers vs payload** | A string is a pointer plus a length; a slice adds a capacity. |
| 3 | **Platform dependence** | Word size makes these numbers architecture-specific. |

## Hint

You need a value to take `Sizeof` of. A zero `Header` will do.

## Validate

```bash
make verify
```
