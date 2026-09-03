# Race the Payment Providers

**Level:** middle
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Checkout is routed through several payment providers. Any of them can authorize
the charge, and their latency varies wildly, so the gateway asks all of them at
once and takes the first authorization code that comes back. The two failure
modes to design around: a provider that never answers must not hold checkout
open, and the goroutines racing for the losers must not be stranded once the
caller has already returned.

## Task

Implement `FirstAuthorization` in [providerrace.go](providerrace.go) so that:

1. It returns `ctx.Err()` when the caller's context is already finished, and
   `ErrAllProvidersDeclined` when `providers` is empty.
2. It derives a cancellable context, starts one goroutine per provider, and
   passes the derived context to `authorize`.
3. Results land on a channel buffered to `len(providers)`, so no loser blocks
   on its send after the winner has returned.
4. The first result with a nil error is returned; cancelling the derived
   context on the way out stops every other provider.
5. If every provider fails, it returns `ErrAllProvidersDeclined`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstAuthorization(live ctx, charge, [slow-a ok-b slow-c], authorize)
Output: "ok-b:auth", nil
```

**Example 2:**

```
Input:  FirstAuthorization(live ctx, charge, [no-a no-b], authorize)
Output: "", ErrAllProvidersDeclined
```

**Example 3:**

```
Input:  FirstAuthorization(cancelled ctx, charge, [ok-a], authorize)
Output: "", context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **First-success race** | Fan out to N, return on the first nil error, ignore the rest. |
| 2 | **`defer cancel()`** | The cancel is what turns "ignore the rest" into "stop the rest". |
| 3 | **Buffer sized to the fan-out** | Every loser can deposit its result and exit even though nobody reads it. |
| 4 | **Error aggregation** | Counting the failures is how you know the race is over and lost. |

## Hint

Loop exactly `len(providers)` times over the results channel. Returning inside
that loop on the first success is what makes it a race; `defer cancel()` is
what makes the losers stop.

## Validate

```bash
make verify
```
