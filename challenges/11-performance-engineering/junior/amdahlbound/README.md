# The Ceiling On Every Optimisation

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Amdahl's law is the arithmetic behind "profile before you optimise". Make 2% of the runtime ten times faster and the program gets 1.8% faster — a week of work for a rounding error. The formula tells you that before you start.

## Task

Implement both functions in [amdahlbound.go](amdahlbound.go):

1. `MaxSpeedup` returns `1 / ((1-p) + p/s)`.
2. A `p` outside `[0,1]` or an `s` below `1` returns `1`.
3. `Ceiling` returns `1 / (1-p)`, the limit as `s` grows without bound; `p >= 1` returns `+Inf`.

## Examples

**Example 1:**

```
Input:  MaxSpeedup(0.5, 2)
Output: 1.333...
```

**Example 2:**

```
Input:  MaxSpeedup(0.02, 10)
Output: about 1.018
```

**Example 3:**

```
Input:  Ceiling(0.9)
Output: 10
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The untouched part dominates** | `1-p` never shrinks, so it sets the floor on total runtime. |
| 2 | **Infinite speedup is still bounded** | Even removing the optimised part entirely only reaches `1/(1-p)`. |
| 3 | **Profile first** | The fraction `p` comes from measurement; guessing it is how weeks get wasted. |

## Topics used again

Float arithmetic, `math.Inf`, guards.

## Hint

`Ceiling` is `MaxSpeedup` with `p/s` driven to zero.

## Validate

```bash
make verify
```
