# An Error That Costs Nothing To Return

**Level:** junior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

An input validator on a public endpoint builds its error with `fmt.Errorf` on every rejected request. A burst of bad traffic turns error reporting into the biggest allocator in the process.

## Task

Implement [sentinelerr.go](sentinelerr.go):

1. Return `ErrNegative` for `n < 0` and `ErrTooLarge` for `n > MaxCount`.
2. Return nil otherwise; the boundaries 0 and `MaxCount` are valid.
3. Zero allocations, including on the failing paths.

Replace the stub body in [sentinelerr.go](sentinelerr.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Validate(5)
Output: <nil>
```

**Example 2:**

```
Input:  Validate(-1)
Output: ErrNegative
```

_Explanation:_ The same value every time — comparable with errors.Is.

**Example 3:**

```
Input:  Validate(1001)
Output: ErrTooLarge
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sentinel errors** | One package-level value, created once at init, returned forever. |
| 2 | **errors.Is** | Comparison against a sentinel is what makes callers able to branch. |
| 3 | **fmt.Errorf allocates** | Formatting builds a new string and a new error value on every call. |

## Hint

The two failures are already declared. Return them.

## Validate

```bash
make verify
```
