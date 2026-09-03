# Every Message In The Chain

**Level:** middle
**Topic:** 04-error-handling

## Context

A structured logger emits one field per layer of a failure so each annotation is searchable on its own.

## Task

Implement `Messages` in [chainmsgs.go](chainmsgs.go):

1. Return the message of every error in the chain, outermost first.
2. Return nil for a nil error.
3. Include the root error's message.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Messages(ErrBase)
Output: ["base failure"]
```

**Example 2:**

```
Input:  Messages(fmt.Errorf("a: %w", ErrBase))
Output: ["a: base failure" "base failure"]
```

**Example 3:**

```
Input:  Messages(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Walking a chain** | Collect at each step, then unwrap. |
| 2 | **Nested messages** | Each layer's message already contains the ones below. |
| 3 | **Nil result** | No error means no messages, not an empty slice. |

## Hint

Append before unwrapping, and let `var out []string` supply the nil result for free.

## Validate

```bash
make verify
```
