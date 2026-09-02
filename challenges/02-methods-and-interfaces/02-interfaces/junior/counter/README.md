# Counter Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Several metric sources report a running count. A dashboard sums them without knowing their types.

## Task

Implement the stub(s) in [counter.go](counter.go):

1. Implement `Count` on `*Clicks` — return the stored total.
2. Implement `Count` on `Fixed` — return the constant value.
3. Implement `Total`, which sums every counter.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Total([]Counter{&Clicks{N: 3}, Fixed(2)})
Output: 5
```

**Example 2:**

```
Input:  Total(nil)
Output: 0
```

**Example 3:**

```
Input:  Total([]Counter{Fixed(-1), Fixed(1)})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **One interface, many implementations** | A struct pointer and a defined int both satisfy `Counter`. |
| 2 | **Interface values in a slice** | `[]Counter` can hold mixed dynamic types. |
| 3 | **Accumulator loop** | Reused from language basics: sum with `range`. |

## Hint

`Total` never mentions `Clicks` or `Fixed` — only `Count()`.

## Validate

```bash
make verify
```
