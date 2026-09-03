# Two Runs, One Table

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`benchstat old.txt new.txt` exists because a single benchmark number means nothing on its own. The comparison pairs results by name, computes the change relative to the baseline, and prints the rows in a stable order so two people reading it see the same table.

## Task

Implement `Compare` in [benchcompare.go](benchcompare.go):

1. Emit one row per name present in *both* maps, with the signed percentage change from base to candidate — negative meaning faster.
2. Order the rows by name, so the output does not depend on map iteration order.
3. Skip names whose baseline is zero or negative; empty input gives an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  Compare({A:100 B:200}, {A:80 B:250})
Output: [{A 100 80 -20} {B 200 250 25}]
```

**Example 2:**

```
Input:  Compare({A:100 OnlyBase:5}, {A:100 OnlyNew:5})
Output: [{A 100 100 0}]
```

**Example 3:**

```
Input:  Compare({A:0}, {A:5})
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Compare, never report in isolation** | A ns/op number only means something next to a baseline. |
| 2 | **Sorted output is reproducible output** | Map order is randomised, so any report from a map must sort. |
| 3 | **Unpaired names are not zeros** | A benchmark missing from one side has no comparison, so it has no row. |

## Topics used again

Maps, the comma-ok idiom, `slices.SortFunc`, structs.

## Hint

Collect the paired rows first, then sort by name.

## Validate

```bash
make verify
```
