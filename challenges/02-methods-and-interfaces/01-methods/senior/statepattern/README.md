# State Machine

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A document walks a fixed pipeline: `Draft → Moderation → Published`. `Published`
is terminal. Encoding those transitions in one method means no caller can invent
an illegal jump.

## Task

Implement `Publish` on `*Document` in [statepattern.go](statepattern.go):

1. `Draft` becomes `Moderation`.
2. `Moderation` becomes `Published`.
3. `Published` stays `Published`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Document{State: Draft}.Publish()
Output: State == Moderation
```

**Example 2:**

```
Input:  Publish() again
Output: State == Published
```

**Example 3:**

```
Input:  Publish() a third time
Output: State == Published  (terminal, idempotent)
```

_Explanation:_ the terminal state absorbs further calls instead of wrapping around.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`iota` enum** | `Draft`, `Moderation`, `Published` are 0, 1, 2 of a named `int` type. |
| 2 | **Explicit transition table** | A `switch` per state documents the machine; arithmetic hides it. |
| 3 | **Idempotent terminal state** | The last state must be a no-op, not a wrap-around. |

## Hint

`d.State++` passes the first two assertions and fails the third — it walks past
`Published`. Write the transitions out as a `switch`.

## Validate

```bash
make verify
```
