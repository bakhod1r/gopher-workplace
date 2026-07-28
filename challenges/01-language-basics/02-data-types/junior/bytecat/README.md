# Strings and Bytes

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

A string and a `[]byte` are convertible both ways with a plain type conversion.
Each conversion copies the data.

## Task

Implement `FromBytes(b)` (bytes → string) and `ToBytes(s)` (string → bytes).

## Examples

```go
FromBytes([]byte{'G','o'}) // => "Go"
ToBytes("Go")              // => [71 111]
FromBytes(nil)             // => ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **string([]byte)** | Converts bytes to a string (copy). |
| 2 | **[]byte(string)** | Converts a string to bytes (copy). |
| 3 | **Immutability** | Strings are read-only; the []byte is a mutable copy. |

## Hint

`string(b)` and `[]byte(s)`.

## Validate

```bash
make verify
```
