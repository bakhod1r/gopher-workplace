# Nth Node From End

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

The two-pointer gap technique: advance a lead pointer n nodes ahead, then move
both until lead reaches the end; the trailing pointer is at the target.

## Task

Implement `NthFromEnd` in [nthfromend.go](nthfromend.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NthFromEnd(1->2->3->4->5, 1)
Output: node 5
```

_Explanation:_ 1st from the end is the last node.

**Example 2:**

```
Input:  NthFromEnd(1->2->3->4->5, 3)
Output: node 3
```

**Example 3:**

```
Input:  NthFromEnd(1->2, 5)
Output: nil
```

_Explanation:_ n larger than the length → nil.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Fixed gap** | Lead pointer starts n ahead. |
| 2 | **Move together** | Both advance until lead is nil. |
| 3 | **Bounds** | Return nil if list shorter than n. |

## Hint

Advance `lead` n steps (return nil if it runs off); then move `lead` and `trail` together until `lead == nil`; return `trail`.

## Validate

```bash
make verify
```
