# Defer LIFO Order

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

Deferred calls stack: they execute in reverse order when the function returns.
A named return value can be appended to from within deferred closures.

## Task

Implement `Order` in [deferorder.go](deferorder.go) using three `defer` statements that append 1, 2, then 3 to a named return slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Order()
Output: [3 2 1]
```

**Example 2:**

```
Input:  defers run LIFO
Output: true
```

**Example 3:**

```
Input:  last deferred runs first
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **defer is LIFO** | The last deferred call runs first. |
| 2 | **Named return + defer** | Deferred closures can mutate the named result. |
| 3 | **Scheduling vs running** | `defer` schedules now, runs at return. |

## Hint

Use a named return `out []int`; write `defer func(){ out = append(out,1) }()` for 1,2,3 in that source order — they run 3,2,1.

## Validate

```bash
make verify
```
