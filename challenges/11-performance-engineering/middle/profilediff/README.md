# What Changed Between Two Profiles

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`pprof -base old.pb.gz new.pb.gz` answers the only question that matters after a change: what moved. A function that vanished from the profile is as interesting as one that appeared, so the diff runs over the union of both name sets — an inner join would hide exactly the rows you shipped the change for.

## Task

Implement `Diff` in [profilediff.go](profilediff.go):

1. Cover the union of both maps, treating a missing side as `0`.
2. Drop rows whose delta is zero.
3. Order by absolute delta descending, then by name ascending.

## Examples

**Example 1:**

```
Input:  Diff({a:10}, {a:4})
Output: [{a 10 4 -6}]
```

**Example 2:**

```
Input:  Diff({gone:7 same:5}, {new:9 same:5})
Output: [{new 0 9 9} {gone 7 0 -7}]
```

**Example 3:**

```
Input:  Diff({a:5}, {a:5})
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Union, not intersection** | Appearing and disappearing are the two most informative outcomes. |
| 2 | **Rank by magnitude** | The biggest movers matter whichever direction they moved. |
| 3 | **A zero delta is not a row** | Unchanged functions are noise in a diff. |

## Topics used again

Maps, set union, `slices.SortFunc`, `cmp.Compare`.

## Hint

Build the key set from both maps first; reading a missing key already gives `0`.

## Validate

```bash
make verify
```
