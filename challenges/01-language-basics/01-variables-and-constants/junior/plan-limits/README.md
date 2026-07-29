# Challenge 01 — Plan Rate Limits

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

You maintain the billing service. Each subscription tier gets a different
API rate limit. The `Tier` type and the `Limit` function are stubbed out —
you need to implement them from scratch.

## Task

Implement [plan.go](plan.go) so that:

1. `Free`, `Pro`, `Enterprise` are **distinct, ascending** values (`0, 1, 2`).
2. `Limit` returns the correct requests-per-minute for each tier.
3. An unknown tier falls back to the `Free` allowance.

Do **not** change the function signature or the tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Limit(Free)
Output: 60
```

**Example 2:**

```
Input:  Limit(Pro)
Output: 600
```

**Example 3:**

```
Input:  Limit(Tier(99))
Output: 60
```

_Explanation:_ Unknown tier falls back to Free's limit.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|


## Hint



## Validate

```bash
make verify
```
