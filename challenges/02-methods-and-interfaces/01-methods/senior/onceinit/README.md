# Once Initialization

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An expensive value must be built lazily, but many goroutines may ask for it at
the same moment. `sync.Once` guarantees the initializer runs exactly once and
that every caller sees the finished result.

## Task

Implement `Get` on `*LazyData` in [onceinit.go](onceinit.go):

1. Use `l.once.Do` to call `l.init()` and assign the result to `l.data`.
2. Return `l.data`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Get() once
Output: "safe"  (init called once)
```

**Example 2:**

```
Input:  Get() from 10 goroutines at once
Output: "safe" for all of them; init call count == 1
```

**Example 3:**

```
Input:  Get() again after the first completed
Output: "safe", init not called
```

_Explanation:_ `Once` remembers that it has fired, forever.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`sync.Once`** | `Do(f)` runs `f` at most once per `Once` value, across all goroutines. |
| 2 | **Once blocks until done** | Late callers wait for the first `Do` to return, so they never observe a half-written `data`. |
| 3 | **Non-copyable state** | A `sync.Once` must not be copied — hence the pointer receiver. |

## Hint

`l.once.Do(func() { l.data = l.init() })` then `return l.data`. The assignment
belongs *inside* the closure.

## Validate

```bash
make verify
```
