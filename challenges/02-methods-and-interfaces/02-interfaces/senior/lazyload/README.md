# Lazy Load

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An expensive index was built at startup even when the request path never touched it. It is now built on first use, exactly once.

## Task

Implement the stub(s) in [lazyload.go](lazyload.go):

1. Implement `Get` on `*Lazy`, building the value on the first call and reusing it afterwards.
2. Implement `Built`, which reports whether the value has been built.
3. Constraint: under concurrent first calls the builder must run exactly **once**, and `-race` must be clean.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Get called twice
Output: the builder runs once
```

**Example 2:**

```
Input:  100 concurrent Get calls
Output: still one build, all callers get the same value
```

**Example 3:**

```
Input:  Built() before any Get
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Once** | Exactly-once initialisation, safe under concurrency. |
| 2 | **Lazy initialisation** | Cost is paid only if the value is used. |
| 3 | **Happens-before via Once** | The builder's writes are visible to every later caller. |

## Hint

`sync.Once.Do` both serialises the first call and publishes its writes to later callers.

## Validate

```bash
make verify
```
