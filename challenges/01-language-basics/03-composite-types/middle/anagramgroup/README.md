# Count Anagram Groups

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Grouping anagrams by a canonical key — a real "normalize then group" task.

## Task

Implement `Count(words)` — number of distinct anagram groups.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ["eat","tea","tan","ate","nat","bat"]
Output: 3
```

_Explanation:_ {eat,tea,ate}, {tan,nat}, {bat}

**Example 2:**

```
Input:  ["a","b","c"]
Output: 3
```

_Explanation:_ no anagrams, 3 distinct

**Example 3:**

```
Input:  ["listen","silent"]
Output: 1
```

_Explanation:_ same sorted letters

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
