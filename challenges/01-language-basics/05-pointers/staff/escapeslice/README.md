# Store Distinct Struct Pointers

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _memory-management_

## Context

`it` is declared ONCE, so `&it` is the same address every iteration; all
pointers see the final value. Declare a fresh struct per iteration so each
pointer escapes to its own allocation.

## Task

Fix [escapeslice.go](escapeslice.go) so each pointer holds a distinct value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Items([1 2 3])
Output: pointers to V=1, V=2, V=3
```

**Example 2:**

```
Input:  deref results
Output: 1, 2, 3 (distinct)
```

**Example 3:**

```
Input:  Items([])
Output: empty
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Hoisted struct address** | One `it`, one address. |
| 2 | **Per-iteration allocation** | Declare inside the loop. |
| 3 | **Escape analysis** | Each fresh struct escapes to the heap. |

## Hint

Allocate per iteration: `it := Item{V: v}; out = append(out, &it)` (declared inside the loop).

## Validate

```bash
make verify
```
