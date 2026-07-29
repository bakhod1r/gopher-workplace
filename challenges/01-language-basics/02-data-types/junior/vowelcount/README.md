# Vowel Count

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Counting vowels means walking characters and testing membership — a natural fit
for ranging over runes and a `switch`.

## Task

Implement `Vowels(s)` counting ASCII vowels (a,e,i,o,u, any case). Accented
letters like `é` do not count.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Vowels("hello")
Output: 2
```

_Explanation:_ e and o.

**Example 2:**

```
Input:  Vowels("AEIOU")
Output: 5
```

_Explanation:_ Case-insensitive, all five.

**Example 3:**

```
Input:  Vowels("xyz")
Output: 0
```

_Explanation:_ No vowels.

**Example 4:**

```
Input:  Vowels("")
Output: 0
```

_Explanation:_ Empty string.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
