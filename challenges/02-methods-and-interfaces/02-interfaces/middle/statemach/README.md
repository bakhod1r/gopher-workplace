# State Machine

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An order moves through states, and each state decides which event it accepts and where it leads.

## Task

Implement the stub(s) in [statemach.go](statemach.go):

1. Implement `Next` and `Name` on `Pending`, `Shipped`, and `Delivered`.
2. `Pending` accepts `"ship"` (to `Shipped`) and `"cancel"` (stays put but reports false); `Shipped` accepts `"deliver"`; `Delivered` accepts nothing.
3. Implement `Run`, which applies a sequence of events and returns the final state name plus how many events were accepted.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Pending{}.Next("ship")
Output: Shipped{}, true
```

**Example 2:**

```
Input:  Delivered{}.Next("ship")
Output: Delivered{}, false
```

**Example 3:**

```
Input:  Run(Pending{}, []string{"ship", "deliver"})
Output: "delivered", 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **State as an interface** | Each state is a type; transitions return the next state. |
| 2 | **Total transitions** | An unhandled event returns the same state and false — no panics. |
| 3 | **Interface return values** | Reused: the method returns the interface, not a concrete type. |

## Hint

Return the receiver itself when the event does not apply.

## Validate

```bash
make verify
```
