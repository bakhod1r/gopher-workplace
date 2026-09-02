# Key And Value

**Level:** junior
**Topic:** 04-error-handling

## Context

An environment file holds `KEY=value` lines. A line without a separator, or with an empty key, is malformed.

## Task

Implement `ParsePair` in [parsepair.go](parsepair.go):

1. Split `s` on the first `=` and return the key and value.
2. Return `ErrNoSeparator` when there is no `=`.
3. Return `ErrEmptyKey` when the key side is empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ParsePair("HOST=local")
Output: "HOST", "local", nil
```

**Example 2:**

```
Input:  ParsePair("HOST")
Output: "", "", ErrNoSeparator
```

**Example 3:**

```
Input:  ParsePair("=x")
Output: "", "", ErrEmptyKey
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **strings.Cut** | Cuts at the first separator only. |
| 2 | **Empty value is legal** | `KEY=` sets an empty value on purpose. |
| 3 | **Zero values on failure** | Both strings are empty when an error is returned. |

## Hint

An empty value is valid; an empty key is not. Only one of the two sides is checked.

## Validate

```bash
make verify
```
