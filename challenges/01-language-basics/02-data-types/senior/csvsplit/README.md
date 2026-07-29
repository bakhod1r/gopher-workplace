# Quoted CSV Split

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

An importer splits CSV lines but breaks `a,"b,c",d` into four fields — it splits
on the comma *inside* the quotes. A comma is only a separator when not inside a
quoted field.

## Task

Fix the `case c == ','` guard between the markers in [csvsplit.go](csvsplit.go)
so commas inside quotes are literal.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a,b,c
Output: ["a" "b" "c"]
```

**Example 2:**

```
Input:  a,"b,c",d
Output: ["a" "b,c" "d"]
```

_Explanation:_ comma inside quotes is literal

**Example 3:**

```
Input:  "x""y",z
Output: [x"y z]
```

_Explanation:_ doubled quote -> literal quote

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Stateful scan** | Track `inQuotes` across characters. |
| 2 | **Context-sensitive delimiter** | Comma splits only outside quotes. |
| 3 | **Quote escaping** | `""` inside quotes is a literal `"`. |

## Hint

`case c == ',' && !inQuotes:`.

## Validate

```bash
make verify
```
