# Firewall Allowlist

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

An edge firewall keeps a set of allowed client IPs. Admin goroutines add addresses while every inbound connection checks membership, so the set is read and written at the same time.

## Task

Implement the stubbed functions in [allowlist.go](allowlist.go) so that:

1. `Allow` adds an IP and reports whether it was newly added.
2. `Allowed` reports whether an IP is in the set.
3. `Size` returns how many IPs are allowed.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  a := NewAllowlist(); a.Allow("10.0.0.1")
Output: true
```

**Example 2:**

```
Input:  a.Allow("10.0.0.1"); a.Allow("10.0.0.1")
Output: true, then false
```

**Example 3:**

```
Input:  a := NewAllowlist(); a.Allowed("10.0.0.9")
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Set as a map** | `map[string]struct{}` stores membership with a zero-byte value. |
| 2 | **sync.Mutex** | Check-and-insert is one critical section, so exactly one caller adds a given IP. |
| 3 | **Comma-ok** | `_, ok := set[k]` tests membership without inserting. |

## Hint

Inside one `Lock`, test membership with comma-ok; if absent, insert and return true, otherwise return false.

## Validate

```bash
make verify
go test -race ./...
```
