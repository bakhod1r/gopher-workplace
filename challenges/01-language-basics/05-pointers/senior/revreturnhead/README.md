# Reverse Returns Old Head

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

After reversal, the original `head` is the LAST node (its Next is nil). The new
head is `prev`. Returning `head` yields a one-element list.

## Task

Fix the return in [revreturnhead.go](revreturnhead.go).

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
| 1 | **New head is prev** | The loop leaves prev at the front. |
| 2 | **Old head becomes tail** | Its Next is now nil. |
| 3 | **Return the right node** | `return prev`. |

## Hint

Return the new head: `return prev`.

## Validate

```bash
make verify
```
