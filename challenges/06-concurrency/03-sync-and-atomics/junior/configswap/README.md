# Hot Config Swap

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A service reloads its configuration on SIGHUP while requests keep flowing. Readers must always see a complete, consistent configuration — never one with the new region and the old version number.

## Task

Implement the stubbed functions in [configswap.go](configswap.go) so that:

1. `Store` publishes a new configuration.
2. `Load` returns the current configuration, or the zero `Config` before anything is published.
3. `Version` returns the published version number.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var s Store; s.Load()
Output: Config{}
```

**Example 2:**

```
Input:  var s Store; s.Store(Config{Version: 2, Region: "eu"}); s.Load().Region
Output: "eu"
```

**Example 3:**

```
Input:  s.Store(Config{Version: 2}); s.Store(Config{Version: 3}); s.Version()
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **atomic.Value** | Publishes a whole value at once; readers see the old one or the new one, never a mix. |
| 2 | **nil before first Store** | `Load` returns `nil` until something is stored — guard the type assertion. |
| 3 | **Consistent type** | Every `Store` must pass the *same* concrete type or it panics. |

## Hint

`v := s.v.Load(); if v == nil { return Config{} }; return v.(Config)`.

## Validate

```bash
make verify
go test -race ./...
```
