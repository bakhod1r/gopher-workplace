# append Result Ignored

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`append` returns a new slice header (length grown, maybe relocated). The code
discards that return (`_ = append(...)`), so `out` never grows.

## Task

Fix the line between the markers in
[appendnotcaptured.go](appendnotcaptured.go) to capture `append`'s result.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3]
Output: [2,4,6]
```

**Example 2:**

```
Input:  []
Output: []
```

**Example 3:**

```
Input:  [-1,0,5]
Output: [-2,0,10]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **append returns** | The new slice is the result. |
| 2 | **Header, not array** | Length lives in the header. |
| 3 | **Assign back** | `out = append(out, ...)`. |

## Hint

`out = append(out, x*2)`.

## Validate

```bash
make verify
```
