# Translate

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A game engine needs to move sprites on screen. Each sprite's position is a
`Point` that gets translated by a delta.

## Task

Implement `Translate` on `*Point` in [translate.go](translate.go):

1. Add `dx` to `X` and `dy` to `Y`.
2. Mutate the receiver — pointer receiver required.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  p := Point{1, 2}; p.Translate(3, 4)
Output: p == Point{4, 6}
```

**Example 2:**

```
Input:  p := Point{5, 5}; p.Translate(-2, -3)
Output: p == Point{3, 2}
```

**Example 3:**

```
Input:  p := Point{0, 0}; p.Translate(1, 1)
Output: p == Point{1, 1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | `*Point` so the mutation is visible to the caller. |
| 2 | **Multiple parameters** | Methods can accept more than one argument. |

## Hint

`p.X += dx; p.Y += dy` — two assignments on a pointer receiver.

## Validate

```bash
make verify
```
