# The Clone That Shares Its Tags

**Level:** staff  
**Topic:** 03-generics

## Context

A draft/publish workflow clones a document before letting an editor change it. Edits to the draft's tag list are appearing on the live document before anyone presses publish.

## Task

Fix the single planted bug in [shallowclonebug.go](shallowclonebug.go):

1. Find and fix the single bug so the clone's Tags do not share storage with the original.
2. Every other field must still be copied.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  c := CloneDoc(d); c.Tags[0] = "x"
Output: d.Tags[0] unchanged
```

**Example 2:**

```
Input:  CloneDoc(d).Title
Output: d.Title
```

**Example 3:**

```
Input:  CloneAll(ds)[0].Tags[0] = "x"
Output: ds[0].Tags[0] unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Struct assignment is shallow** | Copying a struct copies each field's *header*; slice fields keep pointing at the same array. |
| 2 | **Backing-array aliasing** | A returned slice that shares storage lets the caller mutate your internals. |
| 3 | **Deep copy is per-field work** | Only the fields that own indirect storage need explicit duplication. |

## Hint

Which fields of `Doc` survive a plain struct assignment as independent data?

## Validate

```bash
make verify
```
