# Size The Map Before Filling It

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A log analyser counts a few million tokens. The counting map starts empty and rehashes a dozen times on the way up, copying every bucket each time.

## Task

Implement [mapprealloc.go](mapprealloc.go):

1. Return the count of each distinct word.
2. Create the map with a size hint so it does not grow from empty.
3. A nil input returns an empty map.

Replace the stub body in [mapprealloc.go](mapprealloc.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Count([]string{"a","b","a"})
Output: map[a:2 b:1]
```

**Example 2:**

```
Input:  Count(nil)
Output: map[]
```

**Example 3:**

```
Input:  Count([]string{"x"})
Output: map[x:1]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map size hints** | `make(map[K]V, n)` allocates enough buckets for n entries up front. |
| 2 | **Rehashing** | Growing a map reallocates the bucket array and moves entries. |
| 3 | **Zero value of a map entry** | `m[w]++` works on a missing key because the value starts at 0. |

## Hint

`make` takes a second argument for maps too.

## Validate

```bash
make verify
```
