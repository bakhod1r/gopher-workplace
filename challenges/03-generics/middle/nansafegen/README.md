# Ordering And NaN

**Level:** middle  
**Topic:** 03-generics

## Context

A sensor occasionally emits NaN. A naive minimum silently returns garbage, because every comparison with NaN is false.

## Task

Implement the stub(s) in [nansafegen.go](nansafegen.go):

1. Implement `MinIgnoringNaN`, skipping NaN values.
2. Return zero and `false` when the slice is empty or holds only NaN.
3. Detect NaN without importing `math`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MinIgnoringNaN([]float64{3, NaN, 1})
Output: 1, true
```

**Example 2:**

```
Input:  MinIgnoringNaN([]float64{NaN})
Output: 0, false
```

**Example 3:**

```
Input:  MinIgnoringNaN([]float64{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **NaN breaks ordering** | `cmp.Ordered` permits floats, but NaN compares false against everything, including itself. |
| 2 | **The `v != v` idiom** | A value that is not equal to itself is NaN — no import required. |
| 3 | **Seed carefully** | Seeding from `s[0]` fails when `s[0]` is NaN, so track "found" separately. |

## Hint

`v != v` is true only for NaN.

## Validate

```bash
make verify
```
