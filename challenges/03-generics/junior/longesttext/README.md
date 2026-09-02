# Longest Text

**Level:** junior  
**Topic:** 03-generics

## Context

A table renderer sizes a column to its widest entry. Entries may be `string` or a named type like `Label`.

## Task

Implement the stub(s) in [longesttext.go](longesttext.go):

1. Implement `Longest`, returning the element with the greatest length and `true`.
2. On a tie keep the earlier element.
3. Return the zero value and `false` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Longest([]string{"a", "bbb"})
Output: "bbb", true
```

**Example 2:**

```
Input:  Longest([]Label{"xx", "yy"})
Output: Label("xx"), true
```

**Example 3:**

```
Input:  Longest([]string{})
Output: "", false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`len` on a type parameter** | `len(v)` compiles when every type in the set supports it. |
| 2 | **The `~` token** | `~int` means "any type whose underlying type is int", so named types like `type Celsius int` are included. |
| 3 | **Zero value of `T`** | `var zero T` names the zero value of an unknown type. |

## Hint

`~string` lets `len` work and lets named string types in.

## Validate

```bash
make verify
```
