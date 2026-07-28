# Escape Sequences

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

TSV output uses a literal tab between fields and a newline at the end. Both are
escape sequences inside an interpreted string.

## Task

Implement `Row(name, value)` returning `name`, a tab, `value`, and a newline.

## Examples

```go
Row("id", "42")   // => "id\t42\n"
Row("a", "b")     // => "a\tb\n"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Interpreted literal** | Double quotes; escapes are processed. |
| 2 | **`\t` and `\n`** | Tab and newline as single characters. |
| 3 | **Concatenation** | `+` joins strings. |

## Hint

`return name + "\t" + value + "\n"`.

## Validate

```bash
make verify
```
