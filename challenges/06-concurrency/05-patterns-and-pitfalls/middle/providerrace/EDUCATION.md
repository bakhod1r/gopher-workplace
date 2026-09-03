# First-Success Race with Cancellation

## Intuition

A race is a fan-out where you stop reading early. That single asymmetry — N
producers, one consumer that leaves after the first good answer — creates the
two classic bugs.

**Stranded senders.** The losers still finish eventually and still send. If the
results channel is unbuffered, each of them blocks forever on `results <- res`
because the consumer has gone home. Buffering to `len(providers)` gives every
producer a slot, so it can deposit and exit no matter when the winner returned.

**Work that never stops.** Buffering fixes the leak but not the waste: a slow
provider is still holding a connection open for a checkout that already
succeeded. That is what the derived context is for. `defer cancel()` fires on
every return path — the winner, the all-declined path, even a panic — and every
`authorize` call watching `ctx.Done()` unwinds.

Counting is the last piece: the race is lost only after *all* N results have
been seen, so the loop runs exactly `len(providers)` times.

## Approach

1. Return `ctx.Err()` for a dead context and `ErrAllProvidersDeclined` for an
   empty provider list.
2. `runCtx, cancel := context.WithCancel(ctx)`, then `defer cancel()`.
3. `results := make(chan attempt, len(providers))`.
4. Start one goroutine per provider; each calls `authorize(runCtx, …)` and
   sends its `{code, err}`.
5. Receive `len(providers)` times; return on the first nil error.
6. After the loop, report the caller's context error if any, else
   `ErrAllProvidersDeclined`.

## Solution

```go
// FirstAuthorization asks every payment provider to authorize the charge at
// the same time and returns the first authorization code that comes back.
// The losing providers are cancelled through a derived context the moment a
// winner is found, so a slow provider cannot hold the checkout open.
//
// It returns ctx.Err() if the caller's context is already finished, and
// ErrAllProvidersDeclined when no provider authorizes.
//
// Examples:
//
//	FirstAuthorization(live ctx, charge, [slow-a ok-b slow-c], auth) => "ok-b:auth"
//	FirstAuthorization(live ctx, charge, [no-a no-b], auth)          => ErrAllProvidersDeclined
//	FirstAuthorization(cancelled ctx, charge, [ok-a], auth)          => context.Canceled
func FirstAuthorization(ctx context.Context, charge Charge, providers []string, authorize func(context.Context, string, Charge) (string, error)) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if len(providers) == 0 {
		return "", ErrAllProvidersDeclined
	}

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	type attempt struct {
		code string
		err  error
	}
	results := make(chan attempt, len(providers))

	for _, provider := range providers {
		go func(provider string) {
			code, err := authorize(runCtx, provider, charge)
			results <- attempt{code: code, err: err}
		}(provider)
	}

	for range providers {
		res := <-results
		if res.err == nil {
			return res.code, nil
		}
	}

	if err := ctx.Err(); err != nil {
		return "", err
	}
	return "", ErrAllProvidersDeclined
}
```

## Walkthrough

- **`[slow-a, ok-b, slow-c]`.** Three goroutines start. `ok-b` returns
  immediately, so its result is the first (and only) one the loop reads; the
  function returns `"ok-b:auth"`, `defer cancel()` fires, and both `slow-`
  providers unblock from `<-ctx.Done()` and deposit their results into the
  buffered channel that nobody will ever read. Nothing leaks.
- **`[no-a, ok-b, no-c]`.** The declines may arrive before the success; the loop
  skips them because `res.err != nil` and keeps receiving until the winner
  appears. Which order they arrive in does not change the answer.
- **`[no-a, no-b]`.** Two results, both errors, loop ends, and the sentinel
  `ErrAllProvidersDeclined` comes back — a distinguishable error the caller can
  test with `errors.Is` rather than one provider's opaque decline.
- **Cancelled caller.** The guard returns before any goroutine is started.

## Pitfalls

- An unbuffered results channel leaks one goroutine per losing provider — per
  checkout. This is the single most common form of this bug in production.
- Forgetting `defer cancel()` means the losers run to completion, and `go vet`
  will flag the lost cancel function.
- `return "", res.err` on the first *failure* turns the race into "whichever
  provider fails first wins", which is exactly backwards.
- Using a `select` over one channel per provider does not scale and reintroduces
  the blocking problem; one shared buffered channel is the pattern.
- Returning a provider's own decline error when everything fails makes the
  caller's error handling depend on which provider answered last.
