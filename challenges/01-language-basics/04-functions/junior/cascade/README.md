# Cascading Permissions

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _control-flow_

## Context

A `switch` with `fallthrough` lets a matched case continue into the next, so higher access levels accumulate the permissions of lower ones.

## Task

Implement `Access` in [cascade.go](cascade.go) so each level adds the permissions below it, comma-joined highest-first.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Access(3)
Output: "admin,write,read"
```

**Example 2:**

```
Input:  Access(2)
Output: "write,read"
```

**Example 3:**

```
Input:  Access(9)
Output: ""
```

_Explanation:_ No exact case matches, so nothing accumulates.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **switch fallthrough** | A matched case continues into the next. |
| 2 | **accumulation** | Build the string across cascading cases. |
| 3 | **default/empty** | Out-of-range levels yield "". |

## Hint

Use `fallthrough` from level 3 down; join with commas.

## Validate

```bash
make verify
```
