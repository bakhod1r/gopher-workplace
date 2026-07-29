# Sum Through Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Reading through a slice of pointers requires skipping nil entries to avoid a
panic.

## Task

Implement `SumPtrs` in [sumptrs.go](sumptrs.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a, b := 1, 2; SumPtrs([]*int{&a, nil, &b})
Output: 3
```

_Explanation:_ Nil entries contribute nothing.

**Example 2:**

```
Input:  SumPtrs(nil)
Output: 0
```

**Example 3:**

```
Input:  a := 5; SumPtrs([]*int{&a, &a})
Output: 10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Read through pointers** | `*p` per element. |
| 2 | **Nil skip** | Guard before deref. |
| 3 | **Accumulate** | Sum the values. |

## Hint

Range; `if p != nil { total += *p }`.

## Validate

```bash
make verify
```
