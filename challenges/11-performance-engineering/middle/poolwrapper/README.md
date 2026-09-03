# A Typed Front Door For `sync.Pool`

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`sync.Pool` returns `any`, has no reset story, and hands back whatever the last user left in it. Wrapping it in a typed struct is how it becomes usable: the wrapper owns the constructor, the assertion and the reset, and callers only see `Get` and `Put`.

## Task

Implement both methods in [poolwrapper.go](poolwrapper.go):

1. `Get` returns a buffer with length `0` and capacity at least `Size`, defaulting to 1024 when `Size` is zero.
2. A buffer that has been through `Put` must come back empty, never carrying the previous caller's bytes.
3. `Put` ignores a nil buffer, and the whole thing must be safe for concurrent use.

## Examples

**Example 1:**

```
Input:  var p Pool; b := p.Get()
Output: len 0, cap at least 1024
```

**Example 2:**

```
Input:  b = append(b, "secret"...); p.Put(b); p.Get()
Output: an empty buffer
```

**Example 3:**

```
Input:  p.Put(nil)
Output: no panic
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`sync.Pool.New`** | The constructor the pool calls when it has nothing to hand out. |
| 2 | **Reset on the way in or out** | Someone must do it, and doing it in the wrapper means callers cannot forget. |
| 3 | **Pools store `any`** | Putting a `[]byte` boxes the slice header; the wrapper hides that from callers. |

## Topics used again

`sync.Pool`, type assertions, methods on pointer receivers, `append`.

## Hint

`p.pool.New` can be set lazily inside `Get`, or the `Get` path can fall back to `make` when the pool is empty.

## Validate

```bash
make verify
```
