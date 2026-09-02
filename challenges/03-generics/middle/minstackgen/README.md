# Stack With Minimum

**Level:** middle  
**Topic:** 03-generics

## Context

A rollback-aware editor shows the smallest value currently on its undo stack, and the display refreshes constantly.

## Task

Implement the stub(s) in [minstackgen.go](minstackgen.go):

1. Implement `Push`, `Pop`, and `Min`.
2. `Min` must run in constant time — no scanning the stack.
3. All three return the zero value and `false` on an empty stack where applicable.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(3); Push(1); Min()
Output: 1, true
```

**Example 2:**

```
Input:  Push(3); Push(1); Pop(); Min()
Output: 3, true
```

**Example 3:**

```
Input:  Min() on empty
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Auxiliary structures** | A parallel stack of running minima buys O(1) queries for O(n) extra space. |
| 2 | **Keeping the stacks aligned** | Every push and pop must touch both stacks, or they drift. |
| 3 | **Space-time trade-off** | This is the canonical example of paying memory for query speed. |

## Hint

Push a minimum on every push — even when it repeats the previous one.

## Validate

```bash
make verify
```
