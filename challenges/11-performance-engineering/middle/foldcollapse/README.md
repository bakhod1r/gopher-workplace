# Flattening Recursion For The Picture

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A recursive descent parser produces stacks two hundred frames deep, all the same function. Rendered honestly that is a flame graph two hundred rows tall carrying one row of information. Collapsing *consecutive* repeats keeps the picture readable without inventing call paths that never happened.

## Task

Implement both functions in [foldcollapse.go](foldcollapse.go):

1. `Collapse` squashes each run of consecutive identical frames to one.
2. A frame that reappears after a different frame is a real second entry and must survive; the input must not be modified.
3. `Depth` returns how many frames the collapse removed.

## Examples

**Example 1:**

```
Input:  Collapse([a b b b c])
Output: [a b c]
```

**Example 2:**

```
Input:  Collapse([a b a b])
Output: [a b a b]
```

**Example 3:**

```
Input:  Depth([a b b b])
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Consecutive is not the same as duplicate** | `a b a` is mutual recursion, a real shape worth seeing. |
| 2 | **Collapsing is presentation, not aggregation** | The values do not change; only the rendered depth does. |
| 3 | **Compare to the last kept frame** | Comparing to the previous *input* frame gives the same answer here, but the kept-frame rule generalises. |

## Topics used again

Slices, `append` with a capacity hint, string comparison.

## Hint

Keep a frame only when it differs from the one you kept last.

## Validate

```bash
make verify
```
