# Adjacency Graph

**Level:** middle  
**Topic:** 03-generics

## Context

A build system models task dependencies. Nodes are task names, and a duplicate dependency must not be recorded twice.

## Task

Implement the stub(s) in [graphgen.go](graphgen.go):

1. Implement `NewGraph`, `AddEdge`, `Neighbors`, and `Degree`.
2. Adding the same edge twice changes nothing.
3. `Neighbors` returns a copy in insertion order; unknown nodes have none.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AddEdge(a,b); Neighbors(a)
Output: [b]
```

**Example 2:**

```
Input:  AddEdge(a,b) twice; Degree(a)
Output: 1
```

**Example 3:**

```
Input:  Neighbors(unknown)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Adjacency lists** | `map[K][]K` is the compact representation for sparse graphs. |
| 2 | **Nil slices work** | Appending to a missing key needs no initialisation. |
| 3 | **Defensive copies** | Handing out internal storage lets callers corrupt the structure. |

## Hint

Scan the existing list before appending — that is the whole duplicate rule.

## Validate

```bash
make verify
```
