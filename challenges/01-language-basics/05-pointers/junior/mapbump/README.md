# Mutate Through Map of Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Map values that are pointers let you mutate shared variables — the map stores
addresses, not copies.

## Task

Implement `BumpAll` in [mapbump.go](mapbump.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a := 1; BumpAll(map[string]*int{"a": &a})
Output: a == 2
```

**Example 2:**

```
Input:  a, b := 1, 2; BumpAll(m{"a":&a,"b":&b})
Output: a == 2, b == 3
```

**Example 3:**

```
Input:  BumpAll(map[string]*int{})
Output: no-op
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map of pointers** | Values are addresses. |
| 2 | **Range values** | `for _, p := range m`. |
| 3 | **Mutate through** | `*p++`. |

## Hint

Range values; `*p++`.

## Validate

```bash
make verify
```
