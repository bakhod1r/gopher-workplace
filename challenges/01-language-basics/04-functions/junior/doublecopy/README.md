# Value Copy Doesn't Leak

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _call-by-value_

## Context

Go passes arguments by value: the function gets its own copy of `price`, so
arithmetic inside never touches the caller's variable.

## Task

Implement `AddTax` in [doublecopy.go](doublecopy.go): return `price + price*rate/100`.

Do **not** change the function signature or the tests.

## Examples

```go
AddTax(200, 10) // => 220
AddTax(100, 0)  // => 100
AddTax(99, 50)  // => 148
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Call by value** | Parameters are copies of the arguments. |
| 2 | **Integer percent** | `price*rate/100` truncates. |
| 3 | **No side effects** | The caller's variable is unchanged. |

## Hint

`return price + price*rate/100`.

## Validate

```bash
make verify
```
