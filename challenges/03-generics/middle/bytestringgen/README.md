# Strings And Bytes

**Level:** middle  
**Topic:** 03-generics

## Context

A protocol sniffer inspects payloads that arrive as decoded strings in tests and as raw byte slices in production.

## Task

Implement the stub(s) in [bytestringgen.go](bytestringgen.go):

1. Implement `HasPrefix` and `Size` for both string-like and byte-slice-like types.
2. Only operations valid for every member of the set are allowed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  HasPrefix("GET /", "GET")
Output: true
```

**Example 2:**

```
Input:  HasPrefix([]byte("POST"), "GET")
Output: false
```

**Example 3:**

```
Input:  Size([]byte{1, 2})
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Inline union constraints** | The set can be written directly in the parameter list when it is used once. |
| 2 | **What survives a mixed set** | `len` and conversion to `string` work for both; `append` and indexing assignment do not. |
| 3 | **Conversion cost** | `string(v)` on a `[]byte` copies; on a string it is free. |

## Hint

`string(v)` is valid for both members — that is what makes the shared implementation possible.

## Validate

```bash
make verify
```
