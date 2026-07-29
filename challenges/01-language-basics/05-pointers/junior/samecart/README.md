# Same Struct Instance

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Comparing struct POINTERS tests instance identity; two structs with equal
fields are still different instances.

## Task

Implement `Same` in [samecart.go](samecart.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := &Cart{}; Same(c, c)
Output: true
```

**Example 2:**

```
Input:  Same(&Cart{}, &Cart{})
Output: false
```

_Explanation:_ Two separate allocations → different addresses.

**Example 3:**

```
Input:  c := &Cart{}; d := c; Same(c, d)
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer identity** | `a == b` compares addresses. |
| 2 | **Instance vs value** | Equal fields, different instances. |
| 3 | **Struct pointers** | `*Cart` addresses. |

## Hint

`return a == b`.

## Validate

```bash
make verify
```
