# The Delete That Half-Unlinks

**Level:** staff  
**Topic:** 03-generics

## Context

An ordered map backed by a hash index over a doubly linked list reports the right keys after most deletes. But `Len` and `len(Keys())` drift apart over a long run, and iterating backwards resurrects entries that were removed.

## Task

Fix the single planted bug in [ordmapunlinkbug.go](ordmapunlinkbug.go):

1. Find and fix the single bug so an unlinked node is detached from both of its neighbours.
2. Deleting the last entry must leave the tail pointing at the new last node.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Set a,b,c; Delete b; Keys()
Output: [a c]
```

**Example 2:**

```
Input:  Set a,b; Delete b; Set c; Keys()
Output: [a c]
```

**Example 3:**

```
Input:  Set a,b,c; Delete b; RevKeys()
Output: [c a]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer surgery** | A doubly linked node has two incoming edges; unlinking must repair both. |
| 2 | **Structural invariants** | Every operation must restore what the type promises about itself. |
| 3 | **Failures that need scale** | A defect that small inputs cannot express is still a defect; test at size. |

## Hint

A node has a `prev` and a `next`. Count how many pointers the delete repairs.

## Validate

```bash
make verify
```
