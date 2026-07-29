# Closures Over Loop Variable

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Since Go 1.22 the loop variable is per-iteration, so closures created in the
loop each capture their own copy — no more the classic shared-variable bug.

## Task

Implement `Handlers` in [perindex.go](perindex.go): build a slice of `n` closures where the i-th returns `i`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Handlers(3)[1]()
Output: 1
```

**Example 2:**

```
Input:  each closure returns its index
Output: true
```

**Example 3:**

```
Input:  Handlers(2)[0]()
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Per-iteration loop var (Go 1.22)** | Each `i` is a fresh variable. |
| 2 | **Capture in closure** | `func() int { return i }`. |
| 3 | **Slice of funcs** | `[]func() int` collects them. |

## Hint

`for i := 0; i < n; i++ { out = append(out, func() int { return i }) }`.

## Validate

```bash
make verify
```
