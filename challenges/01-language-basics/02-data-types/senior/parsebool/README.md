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

## Examples

```go
Parse("on")  // => true, true
Parse("off") // => false, true
Parse("x")   // => false, false
```

## Topics to Master

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
