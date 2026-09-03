# Total Every Int, However Deep

**Level:** senior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A metrics exporter needs the total of every counter in a nested settings tree. The tree gains a level every quarter and the hand-written walker is always one release behind.

## Task

Implement [walkfields.go](walkfields.go):

1. Total every exported int reachable from `v`.
2. Descend into nested structs, pointers, interfaces, slices and arrays.
3. Skip unexported fields; a nil pointer contributes 0.
4. A bare int is its own total; anything else contributes 0.

Replace the stub body in [walkfields.go](walkfields.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  DeepSum(outer{N:1, In:inner{M:2}})
Output: 3
```

**Example 2:**

```
Input:  DeepSum(outer{N:1})
Output: 1
```

_Explanation:_ The nil `Ptr` contributes nothing.

**Example 3:**

```
Input:  DeepSum([]int{1,2,3})
Output: 6
```

_Explanation:_ Slices are walked too.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recursive reflection** | One function per kind, recursing on the contained Values. |
| 2 | **Value.Elem** | Steps through a pointer or an interface to what it holds. |
| 3 | **Nil guards** | `Elem` on a nil pointer yields an invalid Value. |
| 4 | **Export status on the way down** | Every struct level must be filtered, not just the top one. |

## Hint

One switch on the kind, four interesting cases, and a recursive call in each.

## Validate

```bash
make verify
```
