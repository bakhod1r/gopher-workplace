# Protection Proxy

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A proxy stands in front of a real object and speaks the same method. Here it is
a *protection* proxy: it checks the caller's role and only then forwards the
call, so unauthorized callers never reach the worker at all.

## Task

Implement `Do` on `*Proxy` in [proxyobj.go](proxyobj.go):

1. If `p.role` is `"admin"`, return the result of `p.w.Do()`.
2. Otherwise return `"denied"` **without** calling the worker.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Proxy{role: "user"}.Do()
Output: "denied"  (worker call count unchanged)
```

**Example 2:**

```
Input:  Proxy{role: "admin"}.Do()
Output: "done"  (worker call count +1)
```

**Example 3:**

```
Input:  Proxy{role: ""}.Do()
Output: "denied"
```

_Explanation:_ anything that is not exactly `"admin"` is refused.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Same-signature stand-in** | `Proxy.Do` and `Worker.Do` match, so a proxy is drop-in for callers. |
| 2 | **Guarding before delegating** | The test asserts `w.calls`, so an early return is required, not just a discarded result. |
| 3 | **Shared backing object** | Both proxies point at the same `*Worker`; the counter proves who got through. |

## Hint

Return early. Calling `p.w.Do()` and then throwing the result away still bumps
`w.calls` and fails the test.

## Validate

```bash
make verify
```
