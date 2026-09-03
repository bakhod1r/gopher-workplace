# Upper-Case The Bytes You Were Given

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A header normaliser converts each `[]byte` to a string, upper-cases it, and converts back. Two allocations per header, on every request.

## Task

Implement [upperascii.go](upperascii.go):

1. Upper-case the ASCII letters `a`-`z` of `b`, in place.
2. Return `b` itself — the same array the caller passed.
3. No allocations; leave every other byte untouched.

Replace the stub body in [upperascii.go](upperascii.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Upper([]byte("go1 x"))
Output: "GO1 X"
```

_Explanation:_ Digits and spaces are untouched.

**Example 2:**

```
Input:  Upper([]byte("ALREADY"))
Output: "ALREADY"
```

**Example 3:**

```
Input:  Upper(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **[]byte is mutable, string is not** | Working on the byte slice avoids both conversions. |
| 2 | **ASCII arithmetic** | `c - 'a' + 'A'` is the case flip for ASCII letters. |
| 3 | **In-place transforms** | Returning the input makes the call chainable without copying. |

## Hint

`strings.ToUpper` would cost you two conversions. You already hold the bytes.

## Validate

```bash
make verify
```
