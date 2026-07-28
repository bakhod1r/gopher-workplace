# Run-Length Encode

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A simple compression: replace runs of a repeated character with the character
and a count.

## Task

Implement `Encode(s)` producing `char + count` per run.

## Examples

```go
Encode("aaab") // => "a3b1"
Encode("abc")  // => "a1b1c1"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Run detection** | Count while the byte repeats. |
| 2 | **Builder + Itoa** | Assemble char and number. |
| 3 | **Empty input** | Return "". |

## Hint

Walk bytes; count the current run; when it ends, write byte + `strconv.Itoa(n)`.

## Validate

```bash
make verify
```
