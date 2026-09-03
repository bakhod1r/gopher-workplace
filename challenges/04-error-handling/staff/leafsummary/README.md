# Summarise The Leaves

**Level:** staff
**Topic:** 04-error-handling

## Context

A dashboard groups thousands of shard failures by message so an operator sees three distinct causes instead of three thousand lines.

## Task

Implement `Summary` in [leafsummary.go](leafsummary.go):

1. Return a map from leaf message to occurrence count.
2. Count only leaves, not wrappers or join nodes.
3. Return an empty map for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Summary(errors.Join(ErrA, ErrA, ErrB))
Output: {"a": 2, "b": 1}
```

**Example 2:**

```
Input:  Summary(fmt.Errorf("x: %w", ErrA))
Output: {"a": 1}
```

**Example 3:**

```
Input:  Summary(nil)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Aggregating leaves** | Wrappers are noise in a count. |
| 2 | **Recursive counting** | Both unwrap shapes contribute. |
| 3 | **Empty result** | An initialised map, not nil. |

## Hint

A wrapper's message contains its child's — counting wrappers would double-count every failure.

## Validate

```bash
make verify
```
