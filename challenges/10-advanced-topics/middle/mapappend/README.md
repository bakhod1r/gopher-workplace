# The Append That Never Reached The Map

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A grouping helper appends into a map of slices. It compiles, it runs, and every bucket comes out empty.

## Task

Fix the single planted bug in [mapappend.go](mapappend.go):

1. Append `v` to the slice stored under `key`, creating the entry if absent.
2. A nil map must not panic.
3. Fix the single bug so the appended value reaches the map.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Add(m, "a", 1); Add(m, "a", 2)
Output: m["a"] is [1 2]
```

**Example 2:**

```
Input:  Add(m, "new", 7)
Output: the key is created
```

_Explanation:_ A missing key reads as a nil slice, which append handles.

**Example 3:**

```
Input:  Add(nil, "a", 1)
Output: no panic
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map values are not addressable** | `m[k]` is a copy; there is no slot you can append through. |
| 2 | **append returns a new header** | Its result is the only valid slice afterwards. |
| 3 | **The nil slice appends fine** | A missing key needs no special case. |

## Hint

`append` gives you something back. Where does it go?

## Validate

```bash
make verify
```
