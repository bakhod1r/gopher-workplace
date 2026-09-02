# Lazy Value

**Level:** middle  
**Topic:** 03-generics

## Context

A config field is expensive to derive and most requests never read it, so it should be computed on demand and then remembered.

## Task

Implement the stub(s) in [lazygen.go](lazygen.go):

1. Implement `NewLazy`, `Get`, and `Done`.
2. `compute` runs at most once, even when it returns the zero value.
3. This type is not safe for concurrent use.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  l := NewLazy(f); l.Get(); l.Get()
Output: f called once
```

**Example 2:**

```
Input:  Done() before Get()
Output: false
```

**Example 3:**

```
Input:  compute returning 0
Output: still computed once
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deferred work** | The closure carries the computation until someone needs the result. |
| 2 | **Flag over sentinel** | A `done` bool works even when the computed value is the zero value. |
| 3 | **Concurrency belongs elsewhere** | `sync.Once` is the concurrent version — this one is deliberately plain. |

## Hint

A `done` flag, not a comparison with the zero value.

## Validate

```bash
make verify
```
