# Append Result Discarded

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`append` may return a NEW slice header (new pointer/len/cap). Discarding its
return value loses every element; you must assign it back to `out`.

## Task

Fix [appendignored.go](appendignored.go) so squares accumulate.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Squares(4)
Output: [1 4 9 16]
```

**Example 2:**

```
Input:  Squares(1)
Output: [1]
```

**Example 3:**

```
Input:  Squares(0)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **append returns a slice** | The result may differ from the argument. |
| 2 | **Reassign** | `out = append(out, x)`. |
| 3 | **Header vs backing** | len/cap live in the header you must keep. |

## Hint

Assign the result: `out = append(out, i*i)`.

## Validate

```bash
make verify
```
