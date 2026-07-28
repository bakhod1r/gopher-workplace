# Variadic Join

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

A fixed leading parameter can precede a variadic one. Here `sep` is fixed and
`parts` collects the rest.

## Task

Implement `Join` in [joinv.go](joinv.go) placing `sep` between consecutive parts only.

Do **not** change the function signature or the tests.

## Examples

```go
Join(",")            // => ""
Join("-", "a")       // => "a"
Join("/", "a","b","c") // => "a/b/c"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Fixed + variadic params** | A regular parameter may come before `...string`. |
| 2 | **Separator placement** | `sep` goes between parts, not before the first or after the last. |
| 3 | **Empty and single cases** | Zero parts is `""`; one part has no separator. |

## Hint

Build the result by adding `sep` before every part except the first (`i == 0`).

## Validate

```bash
make verify
```
