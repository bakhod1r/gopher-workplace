# The Reset That Resets A Copy

**Level:** staff  
**Topic:** 03-generics

## Context

A connection-pool sweeper is supposed to blank every pooled entry between requests. It compiles, it runs, and the entries keep their old contents.

## Task

Fix the single planted bug in [ptrmethodsetbug.go](ptrmethodsetbug.go):

1. Find and fix the single bug so the elements of the caller's slice are reset.
2. The function must keep working for any `T` whose pointer type has `Reset`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  s = [{3},{4}]; ResetAll(s)
Output: s == [{0},{0}]
```

**Example 2:**

```
Input:  ResetAll on an empty slice
Output: no-op
```

**Example 3:**

```
Input:  ResetAll on a one-element slice
Output: that element is zeroed
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method sets of pointer receivers** | `Reset` on `*T` is not in `T`'s method set, which is exactly why the constraint is written as `*T`. |
| 2 | **Range copies** | `for _, v := range s` binds a fresh copy each iteration; its address is not the element's address. |
| 3 | **Addressable is not the same as aliasing** | `&v` is a perfectly valid pointer — to the wrong memory. |

## Hint

Whose address is `&v`?

## Validate

```bash
make verify
```
