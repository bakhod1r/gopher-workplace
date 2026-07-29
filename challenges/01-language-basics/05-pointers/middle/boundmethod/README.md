# Bind a Pointer Method

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A method value `c.Inc` bound on a pointer receiver captures the pointer, so the
returned function keeps mutating the same counter.

## Task

Implement `Bind` in [boundmethod.go](boundmethod.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  inc := Bind(c); inc(); inc()
Output: c.N == 2
```

_Explanation:_ The method value captures the pointer receiver.

**Example 2:**

```
Input:  inc := Bind(c); inc()
Output: c.N == 1
```

**Example 3:**

```
Input:  c := &Counter{N:5}; Bind(c)()
Output: c.N == 6
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Method value** | `c.Inc` is a callable bound to c. |
| 2 | **Pointer capture** | The bound receiver is the pointer. |
| 3 | **Shared state** | All calls hit the same counter. |

## Hint

`return c.Inc` (the method value bound to the pointer).

## Validate

```bash
make verify
```
