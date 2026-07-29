# JSON Tag Mismatch

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

The API contract is snake_case. `LastName` is tagged `lastName`, so the JSON key
is wrong and clients break.

## Task

Fix the struct tag between the markers in [jsontag.go](jsontag.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  User{First:Ada, Last:Lovelace}
Output: {"first_name":"Ada","last_name":"Lovelace"}
```

**Example 2:**

```
Input:  User{First:Grace, Last:Hopper}
Output: {"first_name":"Grace","last_name":"Hopper"}
```

**Example 3:**

```
Input:  User{}
Output: {"first_name":"","last_name":""}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Struct tags** | `` `json:"name"` `` renames the key. |
| 2 | **Exact key** | The tag string is the literal JSON key. |
| 3 | **encoding/json** | Reads tags via reflection. |

## Hint

`` `json:"last_name"` ``.

## Validate

```bash
make verify
```
