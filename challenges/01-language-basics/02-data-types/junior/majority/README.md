# Boolean Majority

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Three sensors vote; the system trusts the reading only if at least two agree.

## Task

Implement `Majority(a, b, c)` returning true when ≥2 of the three are true.
Use boolean operators only — no counting with ints.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Majority(false,false,false)
Output: false
```

_Explanation:_ No pair true.

**Example 2:**

```
Input:  Majority(true,true,false)
Output: true
```

_Explanation:_ a&&b true.

**Example 3:**

```
Input:  Majority(false,true,true)
Output: true
```

_Explanation:_ b&&c true.

**Example 4:**

```
Input:  Majority(true,true,true)
Output: true
```

_Explanation:_ All pairs true.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Boolean operators** | `&&`, ` |
| 2 | **Pairwise agreement** | Majority = any two agreeing: `(a&&b) |
| 3 | **Short-circuit** | `&&`/` |

## Hint

`(a && b) || (a && c) || (b && c)`.

## Validate

```bash
make verify
```
