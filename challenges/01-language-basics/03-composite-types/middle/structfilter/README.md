# Filter and Project

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Selecting records by a condition and pulling one field — filter + map.

## Task

Implement `ActiveNames(users)` returning names of active users.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [{ann,true},{bob,false},{cid,true}]
Output: ["ann","cid"]
```

**Example 2:**

```
Input:  nil
Output: []
```

**Example 3:**

```
Input:  [{x,false}]
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Filter** | Keep matching structs. |
| 2 | **Project** | Extract one field. |
| 3 | **append** | Build the result slice. |

## Hint

`for _, u := range users { if u.Active { out = append(out, u.Name) } }`.

## Validate

```bash
make verify
```
