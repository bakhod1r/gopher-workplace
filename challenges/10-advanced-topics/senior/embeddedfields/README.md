# Fields That Came From Somewhere Else

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A mapping layer walks struct fields to build column names. Embedding a shared `Base` struct silently produced a column called `Base` instead of `Base.ID`, and the migration ran anyway.

## Task

Implement [embeddedfields.go](embeddedfields.go):

1. Return the dotted path of every exported leaf field, in declaration order.
2. Descend through embedded structs, prefixing with the embedded type's name.
3. Treat a named struct field as a leaf — do not descend into it.
4. Skip unexported fields; return nil for a non-struct.

Replace the stub body in [embeddedfields.go](embeddedfields.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Paths(User{})
Output: [Base.ID Name Extra]
```

_Explanation:_ `Base` is embedded; `Extra` is named.

**Example 2:**

```
Input:  Paths(Deep{})
Output: [User.Base.ID User.Name User.Extra Note]
```

_Explanation:_ Embedding nests.

**Example 3:**

```
Input:  Paths(&User{})
Output: <nil>
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **StructField.Anonymous** | Marks an embedded field — the flag Go uses for promotion. |
| 2 | **Embedding is not inheritance** | The field is still a field; only its name is implicit. |
| 3 | **Recursion on the type** | `reflect.New(f.Type).Elem().Interface()` gives a zero value to recurse on. |
| 4 | **Leaf choice is a design decision** | Descending into every struct would flatten too much. |

## Hint

`f.Anonymous` is the only thing that distinguishes `Base` from `Extra`.

## Validate

```bash
make verify
```
