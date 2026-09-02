# The Snapshot That Keeps Changing

**Level:** staff  
**Topic:** 03-generics

## Context

A document store lets an editor roll back to a checkpoint. The first rollback looks perfect. Every rollback after it returns the state as of the *last* rollback, and the checkpoint slowly rots into whatever was edited afterwards.

## Task

Fix the single planted bug in [versionsnapaliasbug.go](versionsnapaliasbug.go):

1. Find and fix the single bug so restoring gives the live list its own storage.
2. Restoring the same snapshot twice must produce the same result both times.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Append 1,2,3; id=Snapshot(); Set(0,99); Restore(id); Items()
Output: [1 2 3]
```

**Example 2:**

```
Input:  ...then Set(0,77); Restore(id); Items()
Output: [1 2 3]
```

**Example 3:**

```
Input:  Restore(-1)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing-array aliasing** | Two slice headers over one array make every write a shared write. |
| 2 | **Snapshots must be immutable** | A checkpoint that shares storage with the mutable present is not a checkpoint. |
| 3 | **Copy on the way in and on the way out** | `Snapshot` copies; `Restore` has to copy as well. |

## Hint

After a restore, which array do writes to the live list land in?

## Validate

```bash
make verify
```
