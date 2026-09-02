# Resetter

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A test harness returns fixtures to a clean state between runs.

## Task

Implement the stub(s) in [resetter.go](resetter.go):

1. Implement `Reset` on `*Buffer` — empty the data.
2. Implement `Reset` on `*Gauge` — set the value back to zero.
3. Implement `ResetAll`, which resets every element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  b := &Buffer{Data: []string{"x"}}; b.Reset(); len(b.Data)
Output: 0
```

**Example 2:**

```
Input:  g := &Gauge{Value: 9}; g.Reset(); g.Value
Output: 0
```

**Example 3:**

```
Input:  ResetAll([]Resetter{b, g})
Output: both are cleared
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mutating through an interface** | The interface must hold pointers for the writes to stick. |
| 2 | **Zero values** | Reused from language basics: `nil` slice, `0` int. |
| 3 | **Method with no results** | A contract can be pure side effect. |

## Hint

`b.Data = nil` is enough to empty the slice.

## Validate

```bash
make verify
```
