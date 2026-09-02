# Read-Copy-Update

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

For read-mostly state, RCU beats a lock: readers do one atomic pointer load and
never block. Writers never mutate the live object — they copy it, change the
copy, and atomically publish the new pointer. Old readers keep using the old
version until they are done with it.

## Task

Implement `Update` on `*RCU` in [readcopyp.go](readcopyp.go):

1. Take `r.mu` — it serializes *writers only*.
2. Load the current pointer.
3. Build a **new** `Config` with `newData`.
4. Store the new pointer.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  New(); Update("v2")
Output: the published Config holds "v2"
```

**Example 2:**

```
Input:  a reader that loaded the pointer before the update
Output: still sees "v1" — its object was never mutated
```

**Example 3:**

```
Input:  two concurrent Update calls
Output: serialized by the mutex; one of the two values wins
```

_Explanation:_ the mutex orders writers; readers are never involved in it.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Copy, never mutate** | Mutating the loaded `*Config` would be a data race against every reader. |
| 2 | **Atomic publication** | `Store` makes the new version visible in one indivisible step. |
| 3 | **Writer-only lock** | Readers take nothing — that is the point of RCU. |

## Hint

`defer r.mu.Unlock()` right after `Lock`. And build a fresh `&Config{...}` —
if you find yourself writing `cfg.Data = newData`, that is the bug.

## Validate

```bash
make verify
```
