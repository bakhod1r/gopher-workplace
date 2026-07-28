# Sort Structs by Field

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Sorting records by a field with a stable tie-break — the everyday table sort.

## Task

Implement `ByAge(people)` — sorted by Age, then Name; input untouched.

## Examples

```go
ByAge([{bob,30},{amy,25},{cid,30}]) // => [{amy 25} {bob 30} {cid 30}]
```

## Topics to Master

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
