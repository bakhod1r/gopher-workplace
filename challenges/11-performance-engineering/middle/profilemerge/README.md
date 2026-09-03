# Many Machines, One Profile

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Profiling one pod for thirty seconds tells you about one pod. Fleet-wide questions need the profiles merged — which is a sum, but a sum with two invariants: the merged total equals the sum of the totals, and nobody's input map comes back mutated.

## Task

Implement both functions in [profilemerge.go](profilemerge.go):

1. `Merge` sums the per-function values and the sample counts across all profiles.
2. Nil or empty maps contribute nothing, and no input may be modified.
3. Merging nothing gives a profile with an empty, non-nil map; `Total` sums one profile's values.

## Examples

**Example 1:**

```
Input:  Merge([{{a:1 b:5} 3} {{a:2 c:7} 4}])
Output: {{a:3 b:5 c:7} 7}
```

**Example 2:**

```
Input:  Merge([{nil 2} {{a:1} 1}])
Output: {{a:1} 3}
```

**Example 3:**

```
Input:  Merge(nil)
Output: {{} 0}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Merging is a sum with invariants** | Totals must be conserved; a lost sample is a lost hot spot. |
| 2 | **Maps are references** | Accumulating into `profiles[0].Flat` silently edits the caller's profile. |
| 3 | **Sample counts merge too** | Without them the merged profile cannot be normalised back to a rate. |

## Topics used again

Maps as references, structs, `range`.

## Hint

Start from a fresh map, never from one of the inputs.

## Validate

```bash
make verify
```
