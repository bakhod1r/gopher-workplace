# Merge Override Order

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

Config layering: overrides must win over the base. The copies run base-last, so
base overwrites the override on collisions.

## Task

Fix the copy order between the markers in
[mergedirection.go](mergedirection.go) so `over` wins.

## Examples

```go
Merge({a:1,b:2}, {b:20,c:3}) // => {a:1, b:20, c:3}
```

## Topics to Master

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
