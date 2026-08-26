# Map Slice

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A data export tool needs to convert a list of IDs to strings for JSON output.

## Task

Implement `ToString` on `IntList` in [mapslice.go](mapslice.go):

1. Map each `int` to a `string` (e.g. using `fmt.Sprint` or `strconv.Itoa`).
2. Return a new `[]string` slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IntList{1, 2, 3}.ToString()
Output: {"1", "2", "3"}
```

**Example 2:**

```
Input:  IntList{-5, 0}.ToString()
Output: {"-5", "0"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map operation** | Transforming one slice type to another. |
| 2 | **Pre-allocation** | `make([]string, len(l))` is more efficient than append. |

## Hint

`res := make([]string, len(l))`, then loop and assign `res[i] = fmt.Sprint(v)`.

## Validate

```bash
make verify
```
