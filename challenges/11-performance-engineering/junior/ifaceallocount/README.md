# Boxing Happens At The Boundary

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Putting a value into an interface may allocate; calling a method through one never does. Knowing which side of that line you are on tells you whether a `[]Shape` loop is a cost centre or free — and a pointer receiver settles the question, because the pointer already points at heap memory.

## Task

Implement both in [ifaceallocount.go](ifaceallocount.go):

1. `Area` returns `W * H`.
2. `TotalArea` sums the areas of the shapes; no shapes sums to `0`.
3. Iterating an already-built `[]Shape` must not allocate.

## Examples

**Example 1:**

```
Input:  (&Rect{2, 3}).Area()
Output: 6
```

**Example 2:**

```
Input:  TotalArea([&Rect{2,3} &Rect{1,1} &Rect{0,9}])
Output: 7
```

**Example 3:**

```
Input:  TotalArea(nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The method set rule** | With a pointer receiver, `*Rect` satisfies `Shape` and `Rect` does not. |
| 2 | **Boxing a pointer is free** | The interface stores the pointer directly; nothing new is allocated. |
| 3 | **Dynamic dispatch is not allocation** | The call is an indirect jump, and it costs no memory at all. |

## Topics used again

Interfaces, pointer receivers, method sets, `range`.

## Hint

Both bodies are one loop or one expression; nothing here needs `new`.

## Validate

```bash
make verify
```
