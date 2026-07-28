# CSV Field Quoting

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

An export endpoint writes CSV. Per RFC 4180 a field needs quoting only if it
contains a comma, quote, or newline — and inner quotes are doubled.

## Task

Implement `Quote(s)` following that rule.

## Examples

```go
Quote("plain")     // => "plain"
Quote("a,b")       // => `"a,b"`
Quote(`say "hi"`)  // => `"say ""hi"""`
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Conditional quoting** | Only quote when a special char is present. |
| 2 | **Escaping** | Double each inner `"`. |
| 3 | **String scanning** | Detect specials with strings.ContainsAny or a loop. |

## Hint

If `strings.ContainsAny(s, ",\"\n")`, wrap in quotes and replace `"` with `""`;
else return `s`.

## Validate

```bash
make verify
```
