# Longest Match That Returns The Shortest

**Level:** staff  
**Topic:** 03-generics

## Context

A router matches request paths against registered prefixes, most specific wins. Once a catch-all `""` route was added, every request started landing on it — and before that, `/a/b` requests were being served by the `/a` handler.

## Task

Fix the single planted bug in [radixlongestbug.go](radixlongestbug.go):

1. Find and fix the single bug so the *longest* matching prefix wins.
2. A path with no matching prefix must still report false.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Add("",0), Add("/a",1), Add("/a/b",2); Longest("/a/b/c")
Output: 2, true
```

**Example 2:**

```
Input:  same table; Longest("/a/z")
Output: 1, true
```

**Example 3:**

```
Input:  Add("/a",1); Longest("/z")
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Longest match, not first match** | Scanning short to long and returning early inverts the priority order. |
| 2 | **Keep the best, do not return it** | A best-so-far search must finish its scan before committing. |
| 3 | **A registered empty prefix** | The empty string is a legitimate catch-all and matches every input. |

## Hint

The loop walks prefixes from shortest to longest. Where does it stop?

## Validate

```bash
make verify
```
