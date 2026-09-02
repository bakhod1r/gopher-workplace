# TODO Context for a Legacy Client

## Intuition

`context.TODO()` and `context.Background()` behave identically at runtime — both are empty, never-cancelled contexts. The difference is intent. `Background()` says "this really is the root". `TODO()` says "I owe this call site a real request context". Reviewers and linters key off that distinction during a migration.

## Approach

1. Return `context.TODO()`.

## Solution

```go
// LegacyClientContext returns the context the payment-gateway SDK wrapper uses
// until the surrounding call chain is migrated to accept a real request
// context. It marks the call site as unfinished plumbing rather than claiming
// to be a legitimate root.
//
// Examples:
//
//	LegacyClientContext().Err()               => nil
//	LegacyClientContext().Done()              => nil
//	LegacyClientContext() == context.TODO()   => true
func LegacyClientContext() context.Context {
	return context.TODO()
}
```

## Walkthrough

- `context.TODO()` returns its own package-level singleton, distinct from the one `Background()` returns.
- Comparing two `context.Context` interface values compares the underlying pointers, so `== context.TODO()` is true and `== context.Background()` is false.

## Pitfalls

- Returning `context.Background()` compiles and behaves identically, but loses the migration signal — the test checks identity.
- `TODO()` is not a "do nothing" context. It still must be threaded down; it only marks unfinished plumbing.
