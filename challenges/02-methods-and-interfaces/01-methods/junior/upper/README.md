# Upper

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A search engine normalises queries to uppercase for case-insensitive matching.
The query is stored as a custom `MyString` type.

## Task

Implement `Upper` on `MyString` in [upper.go](upper.go):

1. Return the string converted to uppercase.
2. Use `strings.ToUpper`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MyString("hello").Upper()
Output: "HELLO"
```

**Example 2:**

```
Input:  MyString("Go Gopher").Upper()
Output: "GO GOPHER"
```

**Example 3:**

```
Input:  MyString("").Upper()
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods on defined types** | `type MyString string` gets its own method. |
| 2 | **Type conversion** | Convert `MyString` → `string` for `strings.ToUpper`. |

## Hint

`return strings.ToUpper(string(s))` — convert to `string` first.

## Validate

```bash
make verify
```
