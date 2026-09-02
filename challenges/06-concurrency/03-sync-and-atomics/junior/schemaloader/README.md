# One-Shot Schema Load

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A GraphQL server parses its schema file the first time a query arrives. Parsing is expensive and can fail; whichever result the first attempt produces — schema or error — is what every later caller must receive, and the file must be parsed only once.

## Task

Implement the stubbed functions in [schemaloader.go](schemaloader.go) so that:

1. `Load` runs the parse function on the first call and caches both the schema and the error.
2. Later calls return the cached pair without parsing again, even after a failure.
3. `Attempts` reports how many times the parse function ran.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  l := NewLoader(func() (string, error) { return "schema", nil }); l.Load()
Output: "schema", nil
```

**Example 2:**

```
Input:  l.Load(); l.Load(); l.Attempts()
Output: 1
```

**Example 3:**

```
Input:  l := NewLoader(failing); l.Load(); l.Load()
Output: the same error both times, Attempts() == 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Once** | Runs the function once; every other caller blocks until it returns. |
| 2 | **Caching errors** | A memoised failure stays a failure — `Once` never retries. |
| 3 | **Multiple cached fields** | Both the value and the error are written inside the `Do` closure. |

## Hint

Inside `once.Do`, assign `l.schema, l.err = l.parse()` and bump `l.attempts`; return the fields afterwards.

## Validate

```bash
make verify
go test -race ./...
```
