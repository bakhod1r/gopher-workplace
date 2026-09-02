# Health Checker

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A status page probes every dependency of the platform on each refresh. A slow
dependency must not delay the others, so all probes run concurrently and the
page reports one healthy/unhealthy flag per service, in the order the services
are configured.

## Task

Implement `CheckAll` in [healthchecker.go](healthchecker.go) so that:

1. Return a `[]bool` the same length as `services`.
2. Element `i` is `true` when `probe(services[i])` returns a code in `[200, 400)`.
3. Run each probe in its own goroutine and join them with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CheckAll([]string{"api", "db"}, probe)
Output: [true false]
```

**Example 2:**

```
Input:  CheckAll([]string{"api"}, probe)
Output: [true]
```

**Example 3:**

```
Input:  CheckAll(nil, probe)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Result type differs from input** | The input is `[]string` and the output `[]bool`; only the length has to match. |

## Hint

Call `probe` inside the goroutine, not on the parent — otherwise the probes run
one at a time and you have gained nothing.

## Validate

```bash
make verify
```
