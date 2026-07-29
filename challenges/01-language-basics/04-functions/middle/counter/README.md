# Closure Counter

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A closure captures variables from its enclosing scope by reference. A counter
keeps its state in a captured local that lives on after the outer function returns.

## Task

Implement `MakeCounter` in [counter.go](counter.go) so each call to the returned function yields the next integer starting at 1.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := MakeCounter(); c(); c()
Output: 1, then 2
```

**Example 2:**

```
Input:  two counters are independent
Output: true
```

**Example 3:**

```
Input:  c := MakeCounter(); c()
Output: 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Closures capture by reference** | The inner func shares the outer local `n`. |
| 2 | **State outlives the call** | `n` escapes to the heap and persists. |
| 3 | **Independent instances** | Each call to MakeCounter gets a fresh `n`. |

## Hint

Declare `n := 0` and return `func() int { n++; return n }`.

## Validate

```bash
make verify
```
