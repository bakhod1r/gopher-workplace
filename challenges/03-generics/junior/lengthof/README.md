# Length Of

**Level:** junior  
**Topic:** 03-generics

## Context

A size guard rejects oversized payloads that arrive either as a decoded string or as a raw byte slice.

## Task

Implement the stub(s) in [lengthof.go](lengthof.go):

1. Implement `Length`, returning the number of bytes in `v`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Length("abc")
Output: 3
```

**Example 2:**

```
Input:  Length([]byte{1, 2})
Output: 2
```

**Example 3:**

```
Input:  Length("")
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mixed-kind type sets** | `~string | ~[]byte` mixes two different kinds; only operations common to both are allowed. |
| 2 | **`len` across a type set** | `len` works because it is defined for every member of the set. |
| 3 | **Declaring a constraint** | A constraint is an interface holding a type set instead of methods. |

## Hint

The body is one call to `len` — the constraint does the work.

## Validate

```bash
make verify
```
