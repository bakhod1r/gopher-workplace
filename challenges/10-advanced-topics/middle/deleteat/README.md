# Remove One Element, Release Its Slot

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A registry removes entries with `append(s[:i], s[i+1:]...)`. The entries are half a kilobyte each, and the removed one stays reachable through the slot the shortening left behind.

## Task

Implement [deleteat.go](deleteat.go):

1. Remove the element at `i`, keeping the order of the rest.
2. Clear the vacated slot at the end before shortening.
3. An out-of-range index returns the slice unchanged; allocate nothing.

Replace the stub body in [deleteat.go](deleteat.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  DeleteAt([]*Item{a,b,c}, 1)
Output: [a c]
```

**Example 2:**

```
Input:  s[2] after DeleteAt(s, 0)
Output: nil
```

_Explanation:_ The slot the shift emptied is cleared.

**Example 3:**

```
Input:  DeleteAt(s, -1)
Output: s unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Shift with copy** | `copy(s[i:], s[i+1:])` moves the tail down one place. |
| 2 | **The last slot is now a duplicate** | After the shift it holds a second reference to a live element. |
| 3 | **Reslicing hides, it does not release** | The pointer past the new length is still in the array. |

## Hint

After the shift, what is sitting at the old last index?

## Validate

```bash
make verify
```
