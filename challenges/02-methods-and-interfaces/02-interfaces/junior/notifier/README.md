# Notifier

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An alerting system delivers the same message through several channels and reports how many succeeded.

## Task

Implement the stub(s) in [notifier.go](notifier.go):

1. Implement `Notify` on `*Email` — record the message and return true.
2. Implement `Notify` on `Broken` — always fail (return false).
3. Implement `Broadcast`, which notifies everyone and returns the number of successes.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  e := &Email{}; e.Notify("hi")
Output: true, e.Sent == ["hi"]
```

**Example 2:**

```
Input:  Broken{}.Notify("hi")
Output: false
```

**Example 3:**

```
Input:  Broadcast([]Notifier{&Email{}, Broken{}}, "hi")
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface with a result** | Delivery reports success without exposing the channel. |
| 2 | **Mixed receivers behind one interface** | A pointer type and a value type coexist in `[]Notifier`. |
| 3 | **Counting successes** | Reused: conditional accumulation. |

## Hint

`Broadcast` must call every notifier, even after a failure.

## Validate

```bash
make verify
```
