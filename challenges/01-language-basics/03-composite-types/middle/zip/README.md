# Zip to Map

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Pairing parallel arrays (column names + values) into a map.

## Task

Implement `Zip(keys, vals)` over the shorter length.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  keys=["a","b","c"], vals=[1,2,3]
Output: {a:1,b:2,c:3}
```

**Example 2:**

```
Input:  keys=["a","b"], vals=[1,2,3]
Output: {a:1,b:2}
```

_Explanation:_ extra val ignored

**Example 3:**

```
Input:  keys=[], vals=[1]
Output: {}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Parallel index** | `keys[i]` with `vals[i]`. |
| 2 | **Min length** | Stop at the shorter slice. |
| 3 | **Build a map** | Insert each pair. |

## Hint

`n := min(len(keys), len(vals)); for i := 0; i < n; i++ { m[keys[i]] = vals[i] }`.

## Validate

```bash
make verify
```
