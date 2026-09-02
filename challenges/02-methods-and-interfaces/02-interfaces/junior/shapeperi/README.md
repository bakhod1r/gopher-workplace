# Shape Perimeter

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A fencing quote needs the perimeter of every plot, whatever its shape.

## Task

Implement the stub(s) in [shapeperi.go](shapeperi.go):

1. Implement `Perimeter` on `Rect`, `Square`, and `Circle` (use `2 * math.Pi * R`).
2. Implement `LongestPerimeter`, which returns the largest perimeter, or 0 for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Rect{W: 2, H: 3}.Perimeter()
Output: 10
```

**Example 2:**

```
Input:  Square{Side: 4}.Perimeter()
Output: 16
```

**Example 3:**

```
Input:  LongestPerimeter([]Shape{Square{Side: 1}, Rect{W: 5, H: 5}})
Output: 20
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Three implementers of one interface** | The set of implementers is open-ended. |
| 2 | **math.Pi** | Reused from standard library basics: constants from `math`. |
| 3 | **Running maximum** | Reused: comparing against the best so far. |

## Hint

A circle's perimeter is its circumference: `2 * math.Pi * c.R`.

## Validate

```bash
make verify
```
