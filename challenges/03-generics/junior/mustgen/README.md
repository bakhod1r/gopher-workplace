# Must

**Level:** junior  
**Topic:** 03-generics

## Context

Package-level setup reads values that must exist. A failure there is a programming error, not a runtime condition.

## Task

Implement the stub(s) in [mustgen.go](mustgen.go):

1. Implement `Must`, returning the value when `ok` is true and panicking otherwise.
2. `Lookup` is provided — do not change it.
3. The panic message must contain `Must`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Must(Lookup(m, "a"))
Output: the stored value
```

**Example 2:**

```
Input:  Must(Lookup(m, "missing"))
Output: panics
```

**Example 3:**

```
Input:  Must(0, true)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Adapting two-value returns** | `Must(f())` works because Go passes a multi-value call straight into a matching parameter list. |
| 2 | **Where panics belong** | Reused judgement: startup and tests, never per-request code. |
| 3 | **Type parameters preserve the value** | The wrapper returns `T`, so nothing needs a type assertion. |

## Hint

`Must(Lookup(m, k))` works because a two-value call can fill a two-parameter list directly.

## Validate

```bash
make verify
```
