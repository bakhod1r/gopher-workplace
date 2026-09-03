# Bounded Matching

**Level:** staff
**Topic:** 04-error-handling

## Context

A gateway matches sentinels against errors from untrusted plugins. A chain thousands deep must not turn a match into a hot loop.

## Task

Implement `Within` in [depthlimit.go](depthlimit.go):

1. Report whether `target` appears within the first `max` links of `err`'s chain.
2. Count the outermost error as link 1.
3. Return false for a non-positive `max` or a nil argument.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Within(fmt.Errorf("x: %w", ErrA), ErrA, 2)
Output: true
```

**Example 2:**

```
Input:  Within(deep, ErrA, 1)
Output: false
```

**Example 3:**

```
Input:  Within(ErrA, ErrA, 0)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded traversal** | Depth caps keep matching predictable. |
| 2 | **Per-level comparison** | Compare, then unwrap. |
| 3 | **Hostile inputs** | Chain length is attacker-controlled. |

## Hint

Compare at each level with `errors.Is` on the single link, not on the whole remaining chain.

## Validate

```bash
make verify
```
