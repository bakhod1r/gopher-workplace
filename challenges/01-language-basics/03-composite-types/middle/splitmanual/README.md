# Split by Byte

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Understanding `strings.Split` by writing it: cut at each separator, emit the
piece between cuts.

## Task

Implement `Split(s, sep)` (single-byte separator), matching `strings.Split`
semantics.

## Examples

```go
Split("a,,c", ',') // => ["a" "" "c"]
Split("", ',')     // => [""]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice between cuts** | Track the segment start. |
| 2 | **Empty fields** | Consecutive seps yield "". |
| 3 | **Final field** | Emit the tail after the loop. |

## Hint

Track `start`; on `s[i]==sep`, append `s[start:i]`, set `start=i+1`; append
`s[start:]` at the end.

## Validate

```bash
make verify
```
