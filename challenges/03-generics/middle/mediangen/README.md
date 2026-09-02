# Median

**Level:** middle  
**Topic:** 03-generics

## Context

A latency dashboard reports the median rather than the mean, so one slow outlier cannot dominate the number.

## Task

Implement the stub(s) in [mediangen.go](mediangen.go):

1. Implement `Median`, returning the middle value and `true`.
2. For an even number of samples average the two middle values.
3. Leave the input untouched; return zero and `false` when empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Median([]float64{3, 1, 2})
Output: 2, true
```

**Example 2:**

```
Input:  Median([]float64{1, 2, 3, 4})
Output: 2.5, true
```

**Example 3:**

```
Input:  Median([]float64{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Clone before sorting** | The caller's sample slice must survive the call. |
| 2 | **Even and odd cases** | `len/2` is the upper middle; the lower one is `len/2 - 1`. |
| 3 | **Float averaging** | The even case needs division, which is why the set excludes integers. |

## Hint

`len(c)/2` is the upper middle index — the lower one is one before it.

## Validate

```bash
make verify
```
