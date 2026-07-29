# Run Once

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A boolean captured in a closure gates a one-time action — the essence of
`sync.Once` without the concurrency.

## Task

Implement `Once` in [oncefn.go](oncefn.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  o := Once(f); o(); o()
Output: f runs only the first time
```

**Example 2:**

```
Input:  second call is a no-op
Output: true
```

**Example 3:**

```
Input:  guarded by a captured flag
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Captured flag** | `done bool` remembers prior execution. |
| 2 | **Guard the call** | Run `f` only when `!done`, then set `done`. |
| 3 | **Closure state** | The flag persists across calls. |

## Hint

Capture `done := false`; return `func(){ if !done { done = true; f() } }`.

## Validate

```bash
make verify
```
