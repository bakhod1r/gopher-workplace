# Split by Byte

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Understanding `strings.Split` by writing it: cut at each separator, emit the
piece between cuts.

## Task

Implement `Split(s, sep)` (single-byte separator), matching `strings.Split`
semantics.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "a,b,c", sep=,
Output: ["a","b","c"]
```

**Example 2:**

```
Input:  "a,,c", sep=,
Output: ["a","","c"]
```

_Explanation:_ empty field between seps

**Example 3:**

```
Input:  ",", sep=,
Output: ["",""]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice between cuts** | Track the segment start. |
| 2 | **Empty fields** | Consecutive seps yield "". |
| 3 | **Final field** | Emit the tail after the loop. |

## Hint

Track `start`; on `s[i]==sep`, append `s[start:i]`, set `start=i+1`; append
`s[start:]` at the end.

## Validate

```bash
make verify
```
