# Initialize the Inner Map

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`m[o][i]++` writes into the inner map `m[o]`, but a missing outer key yields a
**nil** inner map — and writing to it panics. You must create the inner map first.

## Task

Add the inner-map initialization between the markers in
[nestedmapinit.go](nestedmapinit.go).

## Examples

```go
Tally([{a,x},{a,x},{a,y}]) // => {a:{x:2, y:1}}
```

## Topics to Master

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
