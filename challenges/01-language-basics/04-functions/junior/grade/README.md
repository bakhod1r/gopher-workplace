# Letter Grade Ladder

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _control-flow_

## Context

An if / else-if / else ladder maps a numeric score to a letter, testing the highest threshold first so boundaries stay inclusive.

## Task

Implement `Grade` in [grade.go](grade.go) mapping 0–100 to A–F.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Grade(95)
Output: "A"
```

**Example 2:**

```
Input:  Grade(70)
Output: "C"
```

**Example 3:**

```
Input:  Grade(59)
Output: "F"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **if/else-if ladder** | Test thresholds high to low. |
| 2 | **inclusive bounds** | `>= 90` catches exactly 90. |
| 3 | **fallthrough default** | Below 60 is F. |

## Hint

`if score >= 90 { return "A" } ... return "F"`.

## Validate

```bash
make verify
```
