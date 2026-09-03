# Ask A Value What It Is

**Level:** junior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A debug endpoint prints the shape of whatever it is handed. `%T` gives the declared type name, but the handler needs to branch on the underlying shape instead.

## Task

Implement [kindof.go](kindof.go):

1. Return the name of `v`'s kind.
2. A nil interface reports `"invalid"`.
3. A named type reports its underlying kind, not its own name.

Replace the stub body in [kindof.go](kindof.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  KindName(3)
Output: "int"
```

**Example 2:**

```
Input:  KindName([]int{1})
Output: "slice"
```

**Example 3:**

```
Input:  KindName(nil)
Output: "invalid"
```

_Explanation:_ There is no type inside a nil interface.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.ValueOf** | Opens an interface value for inspection at run time. |
| 2 | **Kind vs Type** | `Kind` is the underlying shape; `Type` is the declared name. |
| 3 | **The zero Value** | `reflect.ValueOf(nil)` is invalid, and `Kind()` says so instead of panicking. |

## Hint

One expression: value, kind, string.

## Validate

```bash
make verify
```
