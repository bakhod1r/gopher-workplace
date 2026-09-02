# Safe Assertion

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An event router forwards payloads only to handlers whose expected shape they match.

## Task

Implement the stub(s) in [assertsafe.go](assertsafe.go):

1. Implement `Handle` on `*IntHandler` — accept the payload only when it is an `int`, adding it to `Sum`.
2. Implement `Handle` on `*TextHandler` — accept only `string` payloads, appending to `Seen`.
3. Implement `Dispatch`, which offers the payload to every handler and returns how many accepted it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IntHandler.Handle(3)
Output: true, Sum == 3
```

**Example 2:**

```
Input:  IntHandler.Handle("3")
Output: false
```

**Example 3:**

```
Input:  Dispatch([]Handler{intH, textH}, "x")
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Assertion inside an implementation** | Each handler narrows `any` to its own type. |
| 2 | **Interface + comma-ok** | The interface hides the payload type; the handler recovers it safely. |
| 3 | **Pointer receivers for accumulation** | Reused: state must survive the call. |

## Hint

Reject before mutating: assert first, then update state.

## Validate

```bash
make verify
```
