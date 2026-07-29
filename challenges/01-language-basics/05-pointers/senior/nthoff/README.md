# Gap Off-by-One

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

The gap between lead and trail must be exactly n. Advancing lead `n+1` times
(with `<= n`) makes the gap too large, returning the (n+1)-th from the end.

## Task

Fix the lead-advance loop in [nthoff.go](nthoff.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NthFromEnd(1->2->3->4->5, 1)
Output: node 5
```

**Example 2:**

```
Input:  NthFromEnd(1->2->3->4->5, 2)
Output: node 4
```

**Example 3:**

```
Input:  NthFromEnd(1->2->3->4->5, 5)
Output: node 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Exact gap** | Advance lead exactly n steps. |
| 2 | **Loop bound** | `i < n`, not `i <= n`. |
| 3 | **Two-pointer window** | Gap size determines the answer. |

## Hint

Use `for i := 0; i < n; i++` so the gap is exactly n.

## Validate

```bash
make verify
```
