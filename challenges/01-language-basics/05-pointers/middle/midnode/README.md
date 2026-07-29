# Middle of List

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

The slow/fast (tortoise/hare) technique finds the middle in one pass: advance
slow by one and fast by two; when fast reaches the end, slow is at the middle.

## Task

Implement `Middle` in [midnode.go](midnode.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Middle(1->2->3->4->5)
Output: node 3
```

_Explanation:_ Odd length → exact middle.

**Example 2:**

```
Input:  Middle(1->2->3->4)
Output: node 3
```

_Explanation:_ Even length → second of the two middles.

**Example 3:**

```
Input:  Middle(1)
Output: node 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slow/fast pointers** | fast moves twice as fast. |
| 2 | **Termination** | Stop when fast or fast.Next is nil. |
| 3 | **Even-length rule** | Returns the second middle. |

## Hint

Advance `slow = slow.Next` and `fast = fast.Next.Next` while `fast != nil && fast.Next != nil`; return slow.

## Validate

```bash
make verify
```
