# Abort an Upload

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The media service uploads large files to object storage. When a user presses Cancel in the browser, the handler aborts the upload context and the storage client must be able to report *why* the transfer stopped, so the audit log distinguishes a user abort from a network timeout.

## Task

Implement the exported function(s) in [abortupload.go](abortupload.go) so that:

1. It derives a cancellable context from `context.Background()`.
2. It calls the cancel function, then returns `ctx.Err()`.
3. The returned error is `context.Canceled`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  AbortUpload()
Output: context.Canceled
```

**Example 2:**

```
Input:  errors.Is(AbortUpload(), context.Canceled)
Output: true
```

**Example 3:**

```
Input:  errors.Is(AbortUpload(), context.DeadlineExceeded)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`context.WithCancel`** | Returns a derived context plus the function that cancels it. |
| 2 | **`ctx.Err()`** | nil while alive; `context.Canceled` after the cancel func runs. |
| 3 | **`context.Canceled`** | The sentinel for "someone called cancel", compared with `errors.Is` or `==`. |

## Hint

`ctx, cancel := context.WithCancel(context.Background())`, then `cancel()`, then return `ctx.Err()`.

## Validate

```bash
make verify
```
