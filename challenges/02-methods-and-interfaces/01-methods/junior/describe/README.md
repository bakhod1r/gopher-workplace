# Describe

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A debug inspector shows shape details. Each shape type has its own `Describe`
method — same method name, different types.

## Task

Implement `Describe` on both `Circle` and `Square` in [describe.go](describe.go):

1. `Circle.Describe()` → `"Circle(radius=R)"`.
2. `Square.Describe()` → `"Square(side=S)"`.
3. Use `%g` format for numbers (no trailing zeros).

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Circle{5}.Describe()
Output: "Circle(radius=5)"
```

**Example 2:**

```
Input:  Square{4}.Describe()
Output: "Square(side=4)"
```

**Example 3:**

```
Input:  Circle{3.5}.Describe()
Output: "Circle(radius=3.5)"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods vs functions** | Same name `Describe` on different types — methods are scoped to their receiver type. |
| 2 | **Value receiver** | Read-only string formatting. |
| 3 | **fmt.Sprintf** | `%g` formats floats without trailing zeros. |

## Hint

`fmt.Sprintf("Circle(radius=%g)", c.Radius)` — and similarly for Square.

## Validate

```bash
make verify
```
