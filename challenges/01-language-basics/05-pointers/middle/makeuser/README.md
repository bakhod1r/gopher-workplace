# new vs Composite Literal

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

`&User{...}` allocates and initialises a struct in one step, returning its
address — the idiomatic constructor form (clearer than `new(User)` plus field
assignments).

## Task

Implement `NewUser` in [makeuser.go](makeuser.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewUser("ann", 30)
Output: &User{Name:"ann", Age:30}
```

**Example 2:**

```
Input:  u := NewUser("x", 0); u != nil
Output: true
```

**Example 3:**

```
Input:  NewUser("", -1)
Output: &User{Name:"", Age:-1}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Composite literal pointer** | `&User{Name: name, Age: age}`. |
| 2 | **Allocation** | Each call is a new instance. |
| 3 | **Constructor idiom** | Return the pointer. |

## Hint

`return &User{Name: name, Age: age}`.

## Validate

```bash
make verify
```
