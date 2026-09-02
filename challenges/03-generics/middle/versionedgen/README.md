# Versioned Value

**Level:** middle  
**Topic:** 03-generics

## Context

A settings editor offers undo. Every change is kept until the user steps back through them.

## Task

Implement the stub(s) in [versionedgen.go](versionedgen.go):

1. Implement `Set`, `Get`, `Undo`, and `Versions`.
2. `Undo` reports `false` when there is nothing to undo.
3. Undoing past the first value leaves the value unset again.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Set(1); Set(2); Get()
Output: 2, true
```

**Example 2:**

```
Input:  Set(1); Set(2); Undo(); Get()
Output: 1, true
```

**Example 3:**

```
Input:  Undo() with no history
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **History as a stack** | The current value is simply the top of the history. |
| 2 | **Unset versus zero** | An empty history is distinguishable from a stored zero value. |
| 3 | **Unbounded growth** | Every `Set` retains the old value — a real memory consideration. |

## Hint

The current value is the last history entry; undo is a pop.

## Validate

```bash
make verify
```
