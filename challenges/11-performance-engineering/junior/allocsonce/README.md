# Pay For It Once

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Rebuilding a lookup table on every call is the most common accidental hot spot there is: correct, easy to read, and quadratic in disguise. Building it lazily on first use — and keeping it — turns a per-call cost into a one-off.

## Task

Implement `Lookup` in [allocsonce.go](allocsonce.go):

1. Return the position of `w` in `Words` and whether it was found.
2. Build the lookup map on the first call and reuse it afterwards.
3. A warm `Lookup` must not allocate. Ties go to the first occurrence.

## Examples

**Example 1:**

```
Input:  Index{Words: [a b]}.Lookup("b")
Output: 1, true
```

**Example 2:**

```
Input:  Index{Words: [a b a]}.Lookup("a")
Output: 0, true
```

**Example 3:**

```
Input:  Index{}.Lookup("a")
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lazy initialisation** | The cost moves to the first call, and every later call is a map hit. |
| 2 | **Pointer receiver required** | A value receiver would build the map into a copy and throw it away. |
| 3 | **A nil map reads fine** | The build check is about whether it exists, not whether it is empty. |

## Topics used again

Methods on pointer receivers, maps, the comma-ok idiom.

## Hint

`if ix.byWord == nil { ... }` at the top, then one map read.

## Validate

```bash
make verify
```
