# Group By First Letter

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Bucketing items by a key — the map-of-slices pattern.

## Task

Implement `ByFirst(words)` grouping by first byte; skip empty words.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ["apple","banana","avocado","cherry",""]
Output: {a:[apple,avocado], b:[banana], c:[cherry]}
```

_Explanation:_ empty word skipped

**Example 2:**

```
Input:  ["dog"]
Output: {d:[dog]}
```

**Example 3:**

```
Input:  [""]
Output: {}
```

_Explanation:_ only empty -> empty map

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
