# URL Slug

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A CMS turns article titles into URL slugs: lowercase, alphanumerics kept,
everything else collapsed to single dashes, edges trimmed.

## Task

Implement `Slug(s)` per the rules above.

## Examples

```go
Slug("Hello, World!") // => "hello-world"
Slug("  Go 1.26  ")   // => "go-1-26"
Slug("a---b")         // => "a-b"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Rune classification** | Keep `0-9a-z`, lowercase letters. |
| 2 | **Run collapsing** | Emit one `-` per run of separators. |
| 3 | **Trim edges** | No leading/trailing `-`. |

## Hint

Scan runes; append lowercased alnum; for a separator set a "pending dash" flag
that emits a single `-` before the next alnum.

## Validate

```bash
make verify
```
