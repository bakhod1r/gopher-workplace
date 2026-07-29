# Fast Pointer Missing Guard

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

Advancing `fast = fast.Next.Next` requires BOTH `fast != nil` AND
`fast.Next != nil`; without the second guard, an acyclic list dereferences a nil
fast.Next and panics.

## Task

Fix the loop condition in [cyclenoguard.go](cyclenoguard.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  HasCycle(acyclic list)
Output: false, no panic
```

**Example 2:**

```
Input:  HasCycle(cyclic list)
Output: true
```

**Example 3:**

```
Input:  HasCycle(nil)
Output: false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two-step guard** | Need `fast != nil && fast.Next != nil`. |
| 2 | **Nil dereference** | `nil.Next` panics. |
| 3 | **Loop safety** | Both hops must be valid. |

## Hint

Guard both hops: `for fast != nil && fast.Next != nil`.

## Validate

```bash
make verify
```
