# Swap Two

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _multiple-return_

## Context

Returning two values makes a swap a pure expression — no temporary variable,
no mutation of the caller's data.

## Task

Implement `Swap` in [swap2.go](swap2.go).

Do **not** change the function signature or the tests.

## Examples

```go
Swap(1, 2)   // => 2, 1
Swap(-5, 5)  // => 5, -5
x, y = Swap(x, y) // exchange in place at the call site
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Multiple return** | Return `(b, a)`. |
| 2 | **No temp needed** | The return list does the shuffle. |
| 3 | **Value semantics** | Arguments are copies; the caller is untouched unless it reassigns. |

## Hint

`return b, a`.

## Validate

```bash
make verify
```
