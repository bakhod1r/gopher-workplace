# A Regression Guard You Can Commit

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Timing assertions in CI are flaky and everyone learns to ignore them. Allocation counts are not: they are deterministic, they fail loudly when someone adds a `fmt.Sprintf` to a hot path, and they need no baseline file. Wrap the measurement in a small helper and a "must not allocate" test becomes one line.

## Task

Implement `Check` in [zeroalloccheck.go](zeroalloccheck.go):

1. Measure the average allocations of one call to `f`, rounded to the nearest whole allocation.
2. Report `OK` when the measured count is at most `limit` — exactly at the limit passes.
3. Clamp `runs` to at least `1` and a negative `limit` to `0`, and report the clamped limit back.

## Examples

**Example 1:**

```
Input:  Check(100, 0, func(){})
Output: {Allocs: 0, Limit: 0, OK: true}
```

**Example 2:**

```
Input:  Check(100, 0, one make)
Output: {Allocs: 1, Limit: 0, OK: false}
```

**Example 3:**

```
Input:  Check(0, -5, func(){})
Output: {Allocs: 0, Limit: 0, OK: true}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Allocation counts are deterministic** | Unlike time, they can be asserted in CI without flakiness. |
| 2 | **A budget, not an exact match** | `<= limit` lets an implementation improve without breaking the test. |
| 3 | **Report the measurement** | A failure that prints the count tells you how far off you are. |

## Topics used again

`testing.AllocsPerRun`, `math.Round`, structs, function values.

## Hint

Clamp first, measure second, compare last.

## Validate

```bash
make verify
```
