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

```go
Classify(5)  // => pos, 1
Classify(0)  // => zero, 1
Classify(-3) // => neg, 1
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
