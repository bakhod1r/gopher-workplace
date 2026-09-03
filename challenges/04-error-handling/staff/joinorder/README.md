# Stable Ordering

**Level:** staff
**Topic:** 04-error-handling

## Context

A report must be byte-identical across runs so it can be diffed. The failures arrive in nondeterministic order, so they are sorted before joining.

## Task

Implement `Sorted` in [joinorder.go](joinorder.go):

1. Join the non-nil errors sorted by message, ascending.
2. Keep duplicates, so equal messages appear as many times as they occurred.
3. Return nil when nothing failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Sorted(ErrB, ErrA)
Output: "a\nb"
```

**Example 2:**

```
Input:  Sorted(ErrA, ErrA)
Output: "a\na"
```

**Example 3:**

```
Input:  Sorted()
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deterministic output** | Diffable reports require a total order. |
| 2 | **sort.Slice with a key** | Sorting by the message string. |
| 3 | **Sorting before joining** | The join reflects the sorted slice. |

## Hint

Sort a copy of the collected slice, not the caller's variadic argument backing array.

## Validate

```bash
make verify
```
