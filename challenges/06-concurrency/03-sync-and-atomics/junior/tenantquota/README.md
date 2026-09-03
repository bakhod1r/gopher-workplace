# Per-Tenant API Quota

**Level:** junior
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A multi-tenant API charges each request against the calling tenant's quota. Tenants come and go at runtime, and requests for different tenants arrive in parallel, so the counter map must handle concurrent inserts as well as concurrent increments.

## Task

Implement the stubbed functions in [tenantquota.go](tenantquota.go) so that:

1. `Add` charges `n` units to a tenant and returns the new total, creating the tenant on first use.
2. `Used` returns a tenant's total, or 0 for an unknown tenant.
3. `Tenants` returns every charged tenant, sorted.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var q Quota; q.Add("acme", 3)
Output: 3
```

**Example 2:**

```
Input:  q.Add("acme", 2)
Output: 5
```

**Example 3:**

```
Input:  q.Used("nobody")
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `sync.Map` | Built for keys written once and read many times, or for disjoint key sets across goroutines — no external lock needed. |
| 2 | `LoadOrStore` | Atomically fetches the existing value or installs yours; two racing goroutines end up with the same counter. |
| 3 | Pointer values | Store a `*counter`, not an `int64` — the pointer stays stable while the value inside is bumped atomically. |

## Hint

`LoadOrStore` returns the value that is actually in the map. Store a pointer to a struct holding an `atomic.Int64`, then `Add` on it.

## Validate

```bash
make verify
```
