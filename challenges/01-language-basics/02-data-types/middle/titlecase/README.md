# Title Case

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Title-casing upper-cases each word's first letter and lower-cases the rest —
case conversion on ASCII letters via arithmetic or `unicode` helpers.

## Task

Implement `Title(s)` for space-separated words.

## Examples

```go
Title("hello world") // => "Hello World"
Title("GO is FUN")   // => "Go Is Fun"
Title("a")           // => "A"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Case conversion** | `unicode.ToUpper/ToLower` on runes. |
| 2 | **Word boundary** | Track "start of word" while scanning. |
| 3 | **Building output** | strings.Builder or []rune. |

## Hint

Track a `startOfWord` flag; upper the first letter, lower the others; reset on
space.

## Validate

```bash
make verify
```
