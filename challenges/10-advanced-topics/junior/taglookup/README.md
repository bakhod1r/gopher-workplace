# Read A Struct Tag

**Level:** junior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A serialiser needs to honour the same `json` tags the standard library reads, plus a `db` tag of its own. Both live on the field, and both have to be read at run time.

## Task

Implement [taglookup.go](taglookup.go):

1. Return the value of `key` in `field`'s struct tag, and whether it was present.
2. A missing field, a missing key, or a non-struct all report false.

Replace the stub body in [taglookup.go](taglookup.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Tag(row{}, "ID", "json")
Output: "id", true
```

**Example 2:**

```
Input:  Tag(row{}, "Name", "db")
Output: "", false
```

_Explanation:_ The field has no db tag.

**Example 3:**

```
Input:  Tag(row{}, "Missing", "json")
Output: "", false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **StructTag.Lookup** | Returns the value and whether the key was present — unlike `Get`, which cannot tell empty from absent. |
| 2 | **FieldByName** | Finds a field by name and reports whether it exists. |
| 3 | **Tags are strings** | The conventional `key:"value"` format is parsed by the `StructTag` methods. |

## Hint

`Get` returns "" for both an empty tag and a missing one. You need the other method.

## Validate

```bash
make verify
```
