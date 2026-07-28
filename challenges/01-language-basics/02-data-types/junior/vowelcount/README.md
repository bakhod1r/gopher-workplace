# Vowel Count

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Counting vowels means walking characters and testing membership — a natural fit
for ranging over runes and a `switch`.

## Task

Implement `Vowels(s)` counting ASCII vowels (a,e,i,o,u, any case). Accented
letters like `é` do not count.

## Examples

```go
Vowels("hello") // => 2
Vowels("AEIOU") // => 5
Vowels("café")  // => 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Ranging runes** | `for _, r := range s` gives each character. |
| 2 | **Rune literals** | `'a'`, `'E'` are rune constants to compare against. |
| 3 | **switch/case list** | `case 'a', 'e', 'i', 'o', 'u':` groups matches. |

## Hint

Lowercase the rune (or list both cases) and `switch` on the vowel set.

## Validate

```bash
make verify
```
