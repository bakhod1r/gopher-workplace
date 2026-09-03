# Find By Behaviour

**Level:** staff
**Topic:** 04-error-handling

## Context

Several packages produce errors that can report a retry-after hint. The client wants the first such hint, whoever produced it.

## Task

Implement `DelayOf` in [asinterface.go](asinterface.go):

1. Return the delay from the first error implementing `RetryAfter() int`.
2. Search the whole tree, wrapped and joined alike.
3. Return `0, false` when nothing implements it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DelayOf(&Throttled{5})
Output: 5, true
```

**Example 2:**

```
Input:  DelayOf(errors.Join(ErrOther, &Throttled{3}))
Output: 3, true
```

**Example 3:**

```
Input:  DelayOf(ErrOther)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface targets for As** | `errors.As` accepts a pointer to an interface. |
| 2 | **Cross-package behaviour** | No shared concrete type required. |
| 3 | **Tree search** | `As` already covers joins. |

## Hint

Declare the interface as a local variable and pass its address — `errors.As` fills it with the first error satisfying it.

## Validate

```bash
make verify
```
