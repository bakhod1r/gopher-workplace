# Turning Stacks Into A Graph

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

The pprof graph view is built from arcs, not from stacks: every adjacent pair of frames in every sample is one caller-to-callee edge, weighted by the sample's value. That is what lets you see *which* caller is driving a hot function — the question the top listing cannot answer.

## Task

Implement both functions in [callgraphedge.go](callgraphedge.go):

1. `Edges` emits one arc per adjacent frame pair, summed across samples; a stack of `n` frames yields `n-1` arcs.
2. Order arcs by value descending, then caller ascending, then callee ascending; samples with a non-positive value or under two frames contribute nothing.
3. `CalleesOf` lists the direct callees of one function, ordered by arc value descending then name ascending.

## Examples

**Example 1:**

```
Input:  Edges([{[a b c] 5}])
Output: [{a b 5} {b c 5}]
```

**Example 2:**

```
Input:  Edges([{[rec rec rec] 2}])
Output: [{rec rec 4}]
```

**Example 3:**

```
Input:  CalleesOf([{a b 5} {a c 9}], "a")
Output: [c b]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Edges carry the attribution** | An arc's weight says how much of a callee's cost came through that caller. |
| 2 | **Adjacent pairs, not all pairs** | `a;b;c` has arcs `a→b` and `b→c`, never `a→c`. |
| 3 | **Composite map keys** | A comparable struct key aggregates pairs without string juggling. |

## Topics used again

Struct map keys, `slices.SortFunc`, multi-key comparisons.

## Hint

`map[Edge]int64` will not work; key by a caller-callee struct and keep the value separate.

## Validate

```bash
make verify
```
