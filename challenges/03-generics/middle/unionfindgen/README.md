# Disjoint Set

**Level:** middle  
**Topic:** 03-generics

## Context

A deduplication job merges records that any rule says are the same person, then asks which records ended up in one cluster.

## Task

Implement the stub(s) in [unionfindgen.go](unionfindgen.go):

1. Implement `NewDisjoint`, `Find`, `Union`, and `Connected`.
2. `Find` registers an unseen element as its own set.
3. Compress paths so repeated queries stay fast.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Union(a,b); Connected(a,b)
Output: true
```

**Example 2:**

```
Input:  Connected(a,c)
Output: false
```

**Example 3:**

```
Input:  Union(a,b); Union(b,c); Connected(a,c)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Union-find** | Each set is a tree; the root is the set's identity. |
| 2 | **Path compression** | Pointing every visited node straight at the root flattens future lookups. |
| 3 | **Transitivity for free** | Merging roots makes connectivity transitive without extra bookkeeping. |

## Hint

`Find` should rewrite `d.parent[v]` to the root it discovers.

## Validate

```bash
make verify
```
