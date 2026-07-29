# Switch Init Statement

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

Re-calling `compute(x)` in every case runs the expensive work multiple times. A
switch INIT statement (`switch v := compute(x); { ... }`) evaluates it once and
reuses `v` across the cases.

## Task

Fix [switchinit.go](switchinit.go) so compute runs exactly once.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Classify(5)
Output: pos, 1 call
```

**Example 2:**

```
Input:  Classify(0)
Output: zero, 1 call
```

**Example 3:**

```
Input:  Classify(-3)
Output: neg, 1 call
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Switch init statement** | `switch v := f(); { ... }` scopes v to the switch. |
| 2 | **Compute once** | Reuse `v` in the case conditions. |
| 3 | **Side-effect count** | Repeated calls inflate the count. |

## Hint

Compute once with an init statement: `switch v := compute(x); { case v < 0: ...; case v == 0: ...; default: ... }`.

## Validate

```bash
make verify
```
