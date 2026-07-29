# Reverse Loses the Rest

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

You must save `cur.Next` BEFORE overwriting it. The bug sets `cur.Next = prev`
first, so the following `next := cur.Next` reads `prev`, losing the remainder.

## Task

Fix the statement order in [reverselose.go](reverselose.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Reverse(1->2->3)
Output: 3->2->1
```

**Example 2:**

```
Input:  Reverse(nil)
Output: nil
```

**Example 3:**

```
Input:  Reverse(5)
Output: 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Save before overwrite** | Capture next first. |
| 2 | **Order matters** | Reassigning Next destroys the link. |
| 3 | **Three-pointer reversal** | next, then re-point, then advance. |

## Hint

Save first: `next := cur.Next; cur.Next = prev`.

## Validate

```bash
make verify
```
