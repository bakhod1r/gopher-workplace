# Join with Separator

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The inverse of split: glue parts with a separator, but not before the first.

## Task

Implement `Join(parts, sep)` without `strings.Join`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ["a","b","c"], sep=","
Output: "a,b,c"
```

**Example 2:**

```
Input:  ["x"], sep=","
Output: "x"
```

**Example 3:**

```
Input:  [], sep=","
Output: ""
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **strings.Builder** | Efficient string assembly. |
| 2 | **Separator placement** | Before all but the first. |
| 3 | **Empty case** | No parts → "". |

## Hint

For `i, p := range parts`: if `i > 0` write `sep`, then write `p`.

## Validate

```bash
make verify
```
