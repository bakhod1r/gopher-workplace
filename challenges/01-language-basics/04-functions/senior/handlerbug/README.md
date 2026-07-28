# Capture the Value, Not the Cursor

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Here `i` is declared ONCE outside the loop, so all closures share it; after the
loop `i == len(names)` and every closure indexes out of range or returns the
same element. Give each closure its own captured value.

## Task

Fix [handlerbug.go](handlerbug.go) so closure k returns `names[k]`.

Do **not** change the function signature or the tests.

## Examples

```go
Labelers([]string{"a","b","c"})[1]() // => "b"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shared vs per-iteration variable** | A pre-declared `i` is shared by all closures. |
| 2 | **Capture by value** | Bind the current index/element inside the loop. |
| 3 | **Go 1.22 range** | `for i := range` gives per-iteration vars; a manual shared `i` does not. |

## Hint

Capture per iteration: either `for i := 0; i < len(names); i++ { name := names[i]; ... return name }`, or switch to `for i := range names` (per-iteration `i`).

## Validate

```bash
make verify
```
