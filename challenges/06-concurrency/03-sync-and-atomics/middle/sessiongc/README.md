# Session Expiry Sweeper

**Level:** middle
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

An edge gateway keeps a live session per connected user and runs a sweeper that drops sessions idle past a cutoff. Sessions are touched by request goroutines all the time, so the sweeper cannot lock the whole table while it works.

## Task

Implement the stubbed functions in [sessiongc.go](sessiongc.go) so that:

1. `Touch` records a session's last-seen tick.
2. `LastSeen` returns the tick and whether the session is live.
3. `Expire` deletes every session last seen **strictly before** the cutoff and returns the removed IDs, sorted.
4. `Active` counts live sessions. Sweeping must not block concurrent touches.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  s.Touch("u1", 10); s.LastSeen("u1")
Output: 10, true
```

**Example 2:**

```
Input:  s.Touch("u1",1); s.Touch("u2",9); s.Expire(5)
Output: ["u1"]
```

**Example 3:**

```
Input:  var s Store; s.Active()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `sync.Map.Range` | Walks a snapshot-ish view; entries added during the walk may or may not be visited, and deleting inside it is allowed. |
| 2 | `LoadAndDelete` | Removes and reports in one atomic step, so two sweepers cannot both claim the same expired session. |
| 3 | No global lock | The sweeper touches one key at a time — request goroutines keep writing throughout. |
| 4 | Deterministic output | `Range` order is unspecified; sort before returning. |

## Hint

`Store` for touch, `Load` for lookup, `Range` + `LoadAndDelete` for the sweep, and `sort.Strings` on the way out.

## Validate

```bash
make verify
```
