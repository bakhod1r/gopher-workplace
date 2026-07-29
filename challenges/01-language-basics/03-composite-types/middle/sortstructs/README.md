# Sort Structs by Field

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Sorting records by a field with a stable tie-break — the everyday table sort.

## Task

Implement `ByAge(people)` — sorted by Age, then Name; input untouched.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [{bob,30},{amy,25},{cid,30}]
Output: [{amy,25},{bob,30},{cid,30}]
```

_Explanation:_ Age asc, Name asc on tie

**Example 2:**

```
Input:  [{x,1}]
Output: [{x,1}]
```

**Example 3:**

```
Input:  [{b,5},{a,5}]
Output: [{a,5},{b,5}]
```

_Explanation:_ tie broken by Name

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Copy first** | Don't mutate the caller's slice. |
| 2 | **sort.Slice** | A less-func comparator. |
| 3 | **Compound key** | Age, then Name for ties. |

## Hint

Copy with append; `sort.Slice(c, func(i,j){ if age differs, compare; else name })`.

## Validate

```bash
make verify
```
