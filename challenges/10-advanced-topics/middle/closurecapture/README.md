# What A Closure Drags Onto The Heap

**Level:** middle
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A team bans closures from a hot path after seeing them in an allocation profile, then discovers half the codebase's iterators are closures. Nobody can say which ones actually cost anything.

## Task

Implement [closurecapture.go](closurecapture.go):

1. Return a function that yields `start`, then `start+1`, and so on.
2. Two counters must not share state.
3. At most two allocations per `Counter` call.

Replace the stub body in [closurecapture.go](closurecapture.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  c := Counter(1); c(); c()
Output: 1, then 2
```

**Example 2:**

```
Input:  a, b := Counter(10), Counter(10); a(); a(); b()
Output: 10
```

_Explanation:_ Independent captured state.

**Example 3:**

```
Input:  Counter(-2) called three times
Output: -2, -1, 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closures capture by reference** | The variable itself is shared with the returned function, not a copy of it. |
| 2 | **Escape via capture** | A captured variable that outlives the frame moves to the heap. |
| 3 | **Func values** | A closure is a pointer to code plus a pointer to its captured environment. |

## Hint

The counter's state has to live somewhere after `Counter` returns.

## Validate

```bash
make verify
```
