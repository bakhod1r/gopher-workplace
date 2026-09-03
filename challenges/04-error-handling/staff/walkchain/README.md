# Visit Every Error

**Level:** staff
**Topic:** 04-error-handling

## Context

A tracing exporter emits one span event per layer of a failure and must visit them in a defined order.

## Task

Implement `Walk` in [walkchain.go](walkchain.go):

1. Call `visit` for every error in the tree, node before children.
2. Visit joined branches left to right.
3. Do nothing for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Walk(fmt.Errorf("x: %w", ErrA), f)
Output: 2 visits
```

**Example 2:**

```
Input:  Walk(errors.Join(ErrA, ErrB), f)
Output: 3 visits
```

**Example 3:**

```
Input:  Walk(nil, f)
Output: 0 visits
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pre-order traversal** | The node is visited before its children. |
| 2 | **Callback traversal** | The visitor decides what to do. |
| 3 | **Both unwrap shapes** | Wrap descends, join branches. |

## Hint

The join node itself counts as a visit — three for a join of two leaves.

## Validate

```bash
make verify
```
