# Defer Argument Snapshot

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

`defer f(x)` evaluates `x` immediately, at the defer statement — not at call
time. Later changes to `x` don't affect the deferred argument.

## Task

Implement `Snapshot` in [defersnapshot.go](defersnapshot.go): set `x := 1`; use a named return `r`; `defer` a call passing `x` that assigns it into `r`; then set `x = 100`; return. Result must be 1.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Snapshot()
Output: 1
```

**Example 2:**

```
Input:  deferred arg captured before mutation
Output: true
```

**Example 3:**

```
Input:  later x=100 ignored
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Argument evaluation time** | `defer f(x)` reads `x` now, runs `f` later. |
| 2 | **Named return** | The deferred call writes into `r`. |
| 3 | **Snapshot vs closure** | Passing as an argument snapshots; capturing in a closure body does not. |

## Hint

Do `defer func(v int){ r = v }(x)` after `x:=1`, then `x=100`. The `x` is copied at defer-time as 1.

## Validate

```bash
make verify
```
