# One Type, Both Shapes

**Level:** staff
**Topic:** 04-error-handling

## Context

An aggregate error also has a primary cause. It must expose the members to matching while keeping a single-cause accessor for callers that want the main one.

## Task

Implement `Multi` in [unwrapboth.go](unwrapboth.go):

1. Give `*Multi` an `Error() string` returning the primary's message.
2. Give it an `Unwrap() []error` returning the primary followed by the others.
3. Give it a `Primary() error` accessor.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m.Error()
Output: "a"
```

**Example 2:**

```
Input:  errors.Is(m, ErrB)
Output: true
```

**Example 3:**

```
Input:  m.Primary()
Output: ErrA
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Choosing an unwrap shape** | A type implements one or the other, not both. |
| 2 | **Accessors beside the interface** | Extra methods carry the rest. |
| 3 | **Member ordering** | The primary comes first. |

## Hint

Implementing both `Unwrap() error` and `Unwrap() []error` on one type is not allowed — pick the multi shape and expose the primary another way.

## Validate

```bash
make verify
```
