# Registry

**Level:** junior  
**Topic:** 03-generics

## Context

Plugins register themselves by name at startup. A duplicate name is a bug, so the second registration must be rejected rather than silently winning.

## Task

Implement the stub(s) in [registrygen.go](registrygen.go):

1. Implement `NewRegistry`, `Register`, `Lookup`, and `Len`.
2. `Register` refuses to overwrite an existing key and reports `false` in that case.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Register(a, 1)
Output: true
```

**Example 2:**

```
Input:  Register(a, 1); Register(a, 2)
Output: false, value stays 1
```

**Example 3:**

```
Input:  Lookup(missing)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **First write wins** | Checking before assigning is the difference from an ordinary map. |
| 2 | **Reporting outcomes with a bool** | The caller learns whether the registration took effect. |
| 3 | **Map keys need `comparable`** | A generic type storing a map must constrain its key parameter. |

## Hint

Check with comma-ok before assigning — a plain assignment always overwrites.

## Validate

```bash
make verify
```
