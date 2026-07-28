# Embedded Field Shadowing

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`Entity` embeds `Base` (which has `ID`) but also declares its own `ID`. The outer
`ID` **shadows** the promoted one, so `e.ID` is the wrong field.

## Task

Fix the return between the markers in
[embeddingshadow.go](embeddingshadow.go) to read the embedded `Base.ID`.

## Examples

```go
BaseID(Entity{Base{42}, 7}) // => 42
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Embedding** | Fields promote to the outer type. |
| 2 | **Shadowing** | An outer field hides a promoted one. |
| 3 | **Qualify** | `e.Base.ID` reaches the embedded field. |

## Hint

`return e.Base.ID`.

## Validate

```bash
make verify
```
