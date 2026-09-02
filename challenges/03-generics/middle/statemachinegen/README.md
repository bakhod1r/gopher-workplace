# State Machine

**Level:** middle  
**Topic:** 03-generics

## Context

An order lifecycle must reject impossible transitions — a shipped order cannot go back to draft — and both states and events are domain-specific types.

## Task

Implement the stub(s) in [statemachinegen.go](statemachinegen.go):

1. Implement `NewMachine`, `Allow`, `Fire`, and `State`.
2. `Fire` returns `false` and changes nothing when no transition is defined.
3. Two type parameters: one for states, one for events.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Allow(draft, submit, sent); Fire(submit)
Output: true, state becomes sent
```

**Example 2:**

```
Input:  Fire(unknown)
Output: false, state unchanged
```

**Example 3:**

```
Input:  State() initially
Output: the start state
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nested maps** | `map[S]map[E]S` is the transition table; the inner map needs allocating per state. |
| 2 | **Reads of missing keys** | `m.table[m.state][e]` is safe even when the outer key is absent — a nil map reads fine. |
| 3 | **Rejecting is not failing** | Returning `false` keeps invalid input from corrupting the state. |

## Hint

Reading through a missing outer key is safe; writing through one is not.

## Validate

```bash
make verify
```
