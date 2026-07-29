# Toggle a Bool Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Pointers work for any type, including bool. Flip with `!`.

## Task

Implement `Toggle` in [toggle.go](toggle.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  b := false; Toggle(&b)
Output: b == true
```

**Example 2:**

```
Input:  b := true; Toggle(&b)
Output: b == false
```

**Example 3:**

```
Input:  b := false; Toggle(&b); Toggle(&b)
Output: b == false
```

_Explanation:_ Two toggles return to the start.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bool pointer** | `*bool` aliases a flag. |
| 2 | **Logical not** | `*p = !*p`. |
| 3 | **Toggle** | Two flips restore. |

## Hint

`*p = !*p`.

## Validate

```bash
make verify
```
