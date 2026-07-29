# Initialize the Inner Map

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`m[o][i]++` writes into the inner map `m[o]`, but a missing outer key yields a
**nil** inner map — and writing to it panics. You must create the inner map first.

## Task

Add the inner-map initialization between the markers in
[nestedmapinit.go](nestedmapinit.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [[x,a],[x,a],[x,b]]
Output: {x:{a:2,b:1}}
```

**Example 2:**

```
Input:  [[p,q]]
Output: {p:{q:1}}
```

**Example 3:**

```
Input:  []
Output: {}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nested maps** | Inner maps are separate values. |
| 2 | **Nil inner map** | Missing outer key → nil inner. |
| 3 | **Lazy init** | Create inner on first use. |

## Hint

`if m[o] == nil { m[o] = make(map[string]int) }` before `m[o][i]++`.

## Validate

```bash
make verify
```
