# DNS Resolver

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A service-mesh sidecar warms its address table at start-up by resolving every
configured upstream host. Lookups are independent and each takes a network round
trip, so they all run concurrently. A host that fails to resolve is recorded as
`0.0.0.0` rather than dropped, so the table stays aligned with the
configuration.

## Task

Implement `ResolveAll` in [dnsresolver.go](dnsresolver.go) so that:

1. Return a slice of addresses the same length as `hosts`, in configuration order.
2. When `resolve` returns the empty string, record `"0.0.0.0"` instead.
3. Resolve each host in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ResolveAll([]string{"a.io", "b.io"}, lookup)
Output: [10.0.0.1 10.0.0.2]
```

**Example 2:**

```
Input:  ResolveAll([]string{"ghost.io"}, lookup)
Output: [0.0.0.0]
```

**Example 3:**

```
Input:  ResolveAll(nil, lookup)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Keeping the slot** | Failures fill their slot with a placeholder; dropping them would need shared state and lose the order. |

## Hint

Do the empty-string check inside the goroutine, on its own local variable, then
perform a single write to `out[i]`.

## Validate

```bash
make verify
```
