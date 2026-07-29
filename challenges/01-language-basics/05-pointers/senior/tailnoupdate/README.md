# Append Can't Grow Empty List

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

When the list is empty there is no tail to link to; the function must CREATE and
return the first node. Returning nil drops the appended value.

## Task

Fix [tailnoupdate.go](tailnoupdate.go) so appending to an empty list works.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Append(nil, 1)
Output: 1
```

**Example 2:**

```
Input:  Append(1, 2)
Output: 1->2
```

**Example 3:**

```
Input:  Append(nil, 9)
Output: 9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Empty-list case** | Create the first node. |
| 2 | **Return the new head** | Caller adopts it. |
| 3 | **Non-empty case** | Link at the tail (already correct). |

## Hint

Return a new head node: `return &Node{Val: v}`.

## Validate

```bash
make verify
```
