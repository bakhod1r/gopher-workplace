# Sub-Benchmarks That Contaminate Each Other

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

A size sweep shows beautiful superlinear growth: size 1 is fast, size 10 is much worse, size 100 is catastrophic. Run the sizes in the opposite order and the curve inverts. The code is fine; the sub-benchmarks are sharing a buffer.

## Task

Fix the single planted bug in [subbenchsharedstatebug.go](subbenchsharedstatebug.go):

1. Find and fix the one bug so each `RunSize` starts from a clean buffer.
2. Running the same size twice must give the same result.
3. Results must not depend on the order the sizes are run in.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  RunSize(3) twice
Output: 3 both times
```

**Example 2:**

```
Input:  RunAll([1 10 100]) and RunAll([100 10 1])
Output: [1 10 100] and [100 10 1]
```

**Example 3:**

```
Input:  RunSize(5) then RunSize(0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sub-benchmarks share the enclosing scope** | Anything declared outside `b.Run` is shared by every case. |
| 2 | **Order-dependent results are a tell** | If reversing the sweep changes the numbers, the state is leaking. |
| 3 | **Reset, do not reallocate** | `buf[:0]` keeps the capacity while clearing the contents. |

## Hint

The buffer needs one line before the appends start.

## Validate

```bash
make verify
```
