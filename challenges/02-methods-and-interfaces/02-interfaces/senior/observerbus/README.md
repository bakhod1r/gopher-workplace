# Observer Bus

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An event bus notifies subscribers. A subscriber that unsubscribes during a notification used to corrupt the subscriber list.

## Task

Implement the stub(s) in [observerbus.go](observerbus.go):

1. Implement `Subscribe` (returning an unsubscribe function), `Publish`, and `Count` on `*Bus`.
2. A subscriber must be able to unsubscribe from inside its own handler without breaking the notification in progress.
3. Constraint: `Publish` must iterate over a snapshot, and `-race` must be clean with concurrent subscribe and publish.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  two subscribers, one Publish
Output: both receive the event
```

**Example 2:**

```
Input:  a subscriber unsubscribing inside its handler
Output: no panic; it gets no further events
```

**Example 3:**

```
Input:  the returned unsubscribe called twice
Output: safe, count unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Observer pattern** | Publishers know nothing about subscribers. |
| 2 | **Snapshot iteration** | Mutating the subscriber set during notification is the classic bug. |
| 3 | **Idempotent unsubscribe** | Reused: exactly-once cleanup semantics. |

## Hint

Copy the handler list under the lock, release the lock, then call the handlers.

## Validate

```bash
make verify
```
