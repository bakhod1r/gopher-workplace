# Group By First Letter

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Bucketing items by a key — the map-of-slices pattern.

## Task

Implement `ByFirst(words)` grouping by first byte; skip empty words.

## Examples

```go
ByFirst([]string{"apple","banana","avocado"})
// => {'a':[apple avocado], 'b':[banana]}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **map[key][]T** | Accumulate slices per key. |
| 2 | **append to map value** | `m[k] = append(m[k], v)`. |
| 3 | **Nil-safe append** | Appending to a missing key's nil works. |

## Hint

`m[word[0]] = append(m[word[0]], word)` after `make`.

## Validate

```bash
make verify
```
