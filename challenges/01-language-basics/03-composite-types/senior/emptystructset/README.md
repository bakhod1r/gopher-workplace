# Set Membership Value

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

The set uses `map[int]bool`, but members are stored as `false`. Membership is
tested with `inB[x]` (the value), which is always false — so the intersection is
empty.

## Task

Fix the store between the markers in
[emptystructset.go](emptystructset.go) so membership is true.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=[1,2,3], b=[2,3,4]
Output: [2,3]
```

**Example 2:**

```
Input:  a=[1,1,2], b=[1]
Output: [1]
```

**Example 3:**

```
Input:  a=[1,2], b=[3,4]
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bool set** | Store `true` to mark membership. |
| 2 | **Value vs presence** | `inB[x]` reads the stored bool. |
| 3 | **Alternative** | `map[int]struct{}` + comma-ok. |

## Hint

`inB[x] = true`.

## Validate

```bash
make verify
```
