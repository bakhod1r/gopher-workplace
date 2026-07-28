# Count Anagram Groups

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Grouping anagrams by a canonical key — a real "normalize then group" task.

## Task

Implement `Count(words)` — number of distinct anagram groups.

## Examples

```go
Count([]string{"eat","tea","tan","ate","nat","bat"}) // => 3
Count([]string{"listen","silent"})                    // => 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Canonical key** | Sorted letters identify a group. |
| 2 | **Set of keys** | Distinct keys = group count. |
| 3 | **[]byte sort** | Sort a word's bytes to key it. |

## Hint

Key = sorted bytes of the word; insert keys into a set; return `len(set)`.

## Validate

```bash
make verify
```
