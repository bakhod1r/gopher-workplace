# Mode (Most Common)

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The statistical mode: the most frequent value, tie broken deterministically.

## Task

Implement `Mode(xs)`; empty → ok=false; ties → smaller value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,2,3]
Output: (2,true)
```

**Example 2:**

```
Input:  [1,1,2,2]
Output: (1,true)
```

_Explanation:_ tie -> smaller value

**Example 3:**

```
Input:  nil
Output: (0,false)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Count map** | Tally frequencies. |
| 2 | **Argmax** | Track best count and value. |
| 3 | **Deterministic ties** | Prefer the smaller value. |

## Hint

Count; scan the map tracking `(count, value)`, preferring higher count then
smaller value.

## Validate

```bash
make verify
```
