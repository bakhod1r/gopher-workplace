# Append to Nil

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A defensive nil-check returns early, but appending to a **nil** slice is perfectly
valid — it allocates and returns a one-element slice. The guard drops the first
element.

## Task

Remove the unnecessary guard between the markers in
[appendnilreturn.go](appendnilreturn.go).

## Examples

```go
Add(nil, 5) // => [5]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil is appendable** | `append(nil, x)` works. |
| 2 | **Zero value slice** | nil behaves like empty for append/range/len. |
| 3 | **Over-defensive code** | The check is harmful here. |

## Hint

Delete the `if s == nil { return nil }` block.

## Validate

```bash
make verify
```
