# Dereference Before Nil Check

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

The nil check must come BEFORE any dereference. Reading `head.Val` first panics
when head is nil, no matter what the later check does.

## Task

Fix [derefearly.go](derefearly.go) so a nil head returns the default without panicking.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstOr(nil, 7)
Output: 7
```

**Example 2:**

```
Input:  FirstOr(&Node{Val: 3}, 7)
Output: 3
```

**Example 3:**

```
Input:  FirstOr(&Node{Val: 0}, 7)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Guard before deref** | Check nil first. |
| 2 | **Order of operations** | The panic happens at `head.Val`. |
| 3 | **Early return** | Return def on nil. |

## Hint

Check nil first: `if head == nil { return def }; return head.Val`.

## Validate

```bash
make verify
```
