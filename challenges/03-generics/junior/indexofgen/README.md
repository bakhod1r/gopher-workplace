# Index Of

**Level:** junior  
**Topic:** 03-generics

## Context

A tab bar highlights the active item. It needs the position of that item, not just whether it exists.

## Task

Implement the stub(s) in [indexofgen.go](indexofgen.go):

1. Implement `IndexOf`, returning the index of the first element equal to `v`.
2. Return `-1` when no element matches.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IndexOf([]int{5, 7, 7}, 7)
Output: 1
```

**Example 2:**

```
Input:  IndexOf([]string{"a", "b"}, "a")
Output: 0
```

**Example 3:**

```
Input:  IndexOf([]int{5}, 7)
Output: -1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |
| 2 | **Index from `range`** | The first `range` variable is the index — you need it here, unlike in `Contains`. |
| 3 | **Sentinel returns** | Reused from language basics: `-1` is the conventional "not found" index in Go. |

## Hint

Return inside the loop as soon as you match — that gives you the *first* index for free.

## Validate

```bash
make verify
```
