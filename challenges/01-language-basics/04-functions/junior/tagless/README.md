# Tagless Switch Sign

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _control-flow_

## Context

A tagless `switch` (no expression) evaluates boolean cases in order — a clean if/else-if chain.

## Task

Implement `Classify` in [tagless.go](tagless.go) returning the sign of `n`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Classify(-5)
Output: "negative"
```

**Example 2:**

```
Input:  Classify(0)
Output: "zero"
```

**Example 3:**

```
Input:  Classify(7)
Output: "positive"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **tagless switch** | `switch { case cond: }`. |
| 2 | **ordered cases** | First true case wins. |
| 3 | **sign logic** | negative / zero / positive. |

## Hint

`switch { case n < 0: return "negative"; case n == 0: return "zero"; default: return "positive" }`.

## Validate

```bash
make verify
```
