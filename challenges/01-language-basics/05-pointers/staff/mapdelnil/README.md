# Delete vs Nil in Pointer Map

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Setting `m[id] = nil` keeps the key (with a nil value); the entry still counts
in len and appears in iteration. Use the built-in `delete` to remove the key.

## Task

Fix [mapdelnil.go](mapdelnil.go) to actually remove the key.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Remove(m, 1)
Output: key 1 gone, len decreases
```

**Example 2:**

```
Input:  len(m) after Remove
Output: len-1
```

**Example 3:**

```
Input:  _, ok := m[1]
Output: false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil value vs absent key** | A nil value is still an entry. |
| 2 | **delete built-in** | `delete(m, id)` removes the key. |
| 3 | **len/iteration** | Only delete shrinks the map. |

## Hint

Remove the key: `delete(m, id)`.

## Validate

```bash
make verify
```
