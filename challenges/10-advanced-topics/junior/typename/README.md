# What Type Is In This Interface

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A debug handler logs what it was given. `%T` works in a format string, and the handler needs the name as a value it can compare and store.

## Task

Implement [typename.go](typename.go):

1. Return the name of `v`'s dynamic type.
2. A nil interface reports `"<nil>"`.
3. A named type reports its own name, not its underlying kind.

Replace the stub body in [typename.go](typename.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  TypeName(3)
Output: "int"
```

**Example 2:**

```
Input:  TypeName(&widget{})
Output: "*typename.widget"
```

_Explanation:_ The pointer is part of the type.

**Example 3:**

```
Input:  TypeName(nil)
Output: "<nil>"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.TypeOf** | Returns nil for a nil interface, and a Type otherwise. |
| 2 | **Type.String vs Kind** | `String` names the type; `Kind` names its shape. |
| 3 | **Aliases are not types** | `type alias = int` is another spelling of int. |

## Hint

`TypeOf(nil)` is nil, and calling a method on it panics.

## Validate

```bash
make verify
```
