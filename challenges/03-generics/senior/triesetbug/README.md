# The Trie That Marks The Wrong Node

**Level:** senior  
**Topic:** 03-generics

## Context

An autocomplete index says every prefix is a complete word and nothing that was actually inserted is found.

## Task

Fix the single planted bug in [triesetbug.go](triesetbug.go):

1. Find and fix the single bug so the terminal flag lands on the last node of the sequence.
2. A proper prefix of an inserted sequence must not report as present.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Insert([a,b]); Contains([a,b])
Output: true
```

**Example 2:**

```
Input:  Contains([a])
Output: false
```

**Example 3:**

```
Input:  Contains([])
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Order of operations** | Doing the right steps in the wrong order is still a bug. |
| 2 | **Structural invariants** | Every operation must restore what the type promises about itself. |

## Hint

Which node is `n` when `end` is set?

## Validate

```bash
make verify
```
