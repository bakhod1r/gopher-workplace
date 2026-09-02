# TODO Context for a Legacy Client

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The payment-gateway wrapper is being migrated to accept a `context.Context` on every call. Half the call sites are still reached from code that has no request context yet. Those sites need the standard marker so the migration ticket can grep for them, and so nobody mistakes them for a legitimate process root.

## Task

Implement the exported function(s) in [legacyclient.go](legacyclient.go) so that:

1. It returns `context.TODO()`.
2. It must not return `context.Background()`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  LegacyClientContext().Err()
Output: nil
```

**Example 2:**

```
Input:  LegacyClientContext().Done()
Output: nil
```

**Example 3:**

```
Input:  LegacyClientContext() == context.TODO()
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`context.TODO()`** | A marker for "a context belongs here but is not wired up yet". |
| 2 | **`TODO` vs `Background`** | Identical at runtime; they differ in what they signal to readers and static analysis. |
| 3 | **Empty contexts** | Both are never cancelled and carry no values. |

## Hint

Return `context.TODO()`. It behaves exactly like `Background()` but documents that the call site is unfinished.

## Validate

```bash
make verify
```
