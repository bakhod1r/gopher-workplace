# Health Fan-Out

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

A status dashboard probes every registered service on each refresh. The probes are independent and slow, so they run in parallel — but the resulting incident list is rendered to humans and paged out, so it must come back in a stable order rather than in whichever order the goroutines happened to finish.

## Task

Implement the exported function(s) in [healthfanout.go](healthfanout.go) so that:

1. Probe every service in its own goroutine, joined with a `sync.WaitGroup`.
2. Collect the names for which `probe` returned a non-nil error.
3. Guard the shared result slice with a `sync.Mutex` — `append` from several goroutines is a data race.
4. Return the collected names sorted alphabetically.
5. Return an empty non-nil slice when nothing is unhealthy or there are no services.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  UnhealthyServices([]string{"api", "db"}, probe)   // db refuses
Output: [db]
```

**Example 2:**

```
Input:  UnhealthyServices([]string{"zeta", "alpha"}, probe)  // both refuse
Output: [alpha zeta]
```

**Example 3:**

```
Input:  UnhealthyServices([]string{"api"}, probe)  // healthy
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mutex-guarded aggregation** | When the result is a growing slice rather than fixed slots, the append must be locked. |
| 2 | **Sorting for determinism** | Goroutine completion order is arbitrary; sorting is what makes the output reproducible. |
| 3 | **`sync.WaitGroup`** | Sort only after `Wait` returns, otherwise a late goroutine appends behind you. |
| 4 | **Nil vs empty slice** | `var s []string` is nil; the tests expect `[]string{}` so the JSON payload stays an array. |

## Hint

Lock only around the `append`, never around the `probe` call itself — holding the mutex across the slow work would serialise the whole fan-out.

## Validate

```bash
make verify
```
