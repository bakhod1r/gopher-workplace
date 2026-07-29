# Config Bool Parsing

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A config loader accepts `true/false`, `1/0`, `yes/no`, `on/off`. But `off` is
not recognized — it falls through to "unknown" and the feature flag defaults
wrong.

## Task

Fix the falsey `case` between the markers in [parsebool.go](parsebool.go) to
include `"off"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "TRUE"
Output: true, true
```

**Example 2:**

```
Input:  "off"
Output: false, true
```

_Explanation:_ must be recognized

**Example 3:**

```
Input:  "maybe"
Output: false, false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Case list** | `case "false", "0", "no", "off":` groups values. |
| 2 | **Normalize input** | Lowercase + trim before matching. |
| 3 | **(value, ok)** | Distinguish "false" from "unrecognized". |

## Hint

Add `"off"` to the falsey case.

## Validate

```bash
make verify
```
