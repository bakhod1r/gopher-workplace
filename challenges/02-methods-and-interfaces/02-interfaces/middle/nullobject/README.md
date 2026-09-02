# Null Object

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Optional collaborators are everywhere in this codebase, and every call site is guarded by a nil check. A do-nothing implementation removes them all.

## Task

Implement the stub(s) in [nullobject.go](nullobject.go):

1. Implement `Report` on `*Recorder` (record the metric) and on `NopMetrics` (do nothing).
2. Implement `MetricsOr`, which returns the given metrics or `NopMetrics{}` when it is nil.
3. Implement `Process`, which reports one metric per item and returns the item count — with no nil check in its body.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Process(rec, []string{"a"}) with a Recorder
Output: 1, rec holds ["item:a"]
```

**Example 2:**

```
Input:  Process(nil, []string{"a"})
Output: 1, nothing recorded
```

**Example 3:**

```
Input:  MetricsOr(nil)
Output: NopMetrics{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Null object pattern** | A do-nothing implementation replaces scattered nil checks. |
| 2 | **Nil interface handling** | The check happens once, at the boundary. |
| 3 | **Interface as a parameter** | Reused: the callee never learns the concrete type. |

## Hint

Normalise at the entry point: `m = MetricsOr(m)` once, then call freely.

## Validate

```bash
make verify
```
