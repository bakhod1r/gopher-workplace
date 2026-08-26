# Encoder

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A report system encodes records as CSV. Each call to `Encode` adds one row.

## Task

Implement `Encode` on `*CSVEncoder` in [encoder.go](encoder.go):

1. Join `fields` with commas.
2. Append the joined string to `Rows`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  e.Encode("a", "b", "c")
Output: Rows appended with "a,b,c"
```

**Example 2:**

```
Input:  e.Encode("only")
Output: Rows appended with "only"
```

**Example 3:**

```
Input:  e.Encode()
Output: Rows appended with ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic methods** | `fields ...string` accepts 0+ args. |
| 2 | **Pointer receiver** | Appending to `Rows` must persist. |
| 3 | **strings.Join** | Join a slice with a separator. |

## Hint

`strings.Join(fields, ",")` then `append` to `e.Rows`.

## Validate

```bash
make verify
```
