# Merge Override Order

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

Config layering: overrides must win over the base. The copies run base-last, so
base overwrites the override on collisions.

## Task

Fix the copy order between the markers in
[mergedirection.go](mergedirection.go) so `over` wins.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  base={a:1,b:2}, over={b:9}
Output: {a:1,b:9}
```

**Example 2:**

```
Input:  base={x:1}, over={y:2}
Output: {x:1,y:2}
```

**Example 3:**

```
Input:  base={k:1}, over={k:1}
Output: {k:1}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Last write wins** | The final assignment to a key stands. |
| 2 | **Layer order** | Base first, override second. |
| 3 | **Map copy** | Iterate and assign. |

## Hint

Copy `base` first, then `over`.

## Validate

```bash
make verify
```
