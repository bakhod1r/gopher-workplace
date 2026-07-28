# Boolean Majority

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Three sensors vote; the system trusts the reading only if at least two agree.

## Task

Implement `Majority(a, b, c)` returning true when ≥2 of the three are true.
Use boolean operators only — no counting with ints.

## Examples

```go
Majority(true, true, false)  // => true
Majority(true, false, false) // => false
Majority(false, true, true)  // => true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Boolean operators** | `&&`, `||` combine truth values. |
| 2 | **Pairwise agreement** | Majority = any two agreeing: `(a&&b)||(a&&c)||(b&&c)`. |
| 3 | **Short-circuit** | `&&`/`||` stop early once the result is known. |

## Hint

`(a && b) || (a && c) || (b && c)`.

## Validate

```bash
make verify
```
