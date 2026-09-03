# Create The Map A Struct Field Needs

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A metadata helper writes into a struct's `Tags` map. Half the structs it is given have never had a tag set, and the helper panics on the first one.

## Task

Implement [mapfield.go](mapfield.go):

1. Set `Tags[key] = val` on the struct `ptr` points at.
2. Create the map when the field is nil; reuse it when it is not.
3. Return `ErrTarget` for a bad target: not a pointer, not a struct, no `Tags` field, unexported, wrong kind, or wrong key/value types.

Replace the stub body in [mapfield.go](mapfield.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  PutTag(&doc{}, "a", "1")
Output: nil, the map is created
```

**Example 2:**

```
Input:  a doc that already has tags
Output: the existing map is kept
```

**Example 3:**

```
Input:  &struct{Tags []string}{}
Output: ErrTarget
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil maps are read-only** | Writing to one panics, in reflection as in ordinary code. |
| 2 | **reflect.MakeMap** | Builds a map of a type known only at run time. |
| 3 | **SetMapIndex** | The reflective map write; the key and value must be assignable to the map's types. |
| 4 | **Type checks before writes** | Every reflective write panics on a mismatch, so validate first. |

## Hint

Check `f.IsNil()` before writing, and `MakeMap` needs the field's own type.

## Validate

```bash
make verify
```
