# Chain of Responsibility

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A request travels down a chain of handlers. Each link either answers it or
passes it along. `BaseHandler` supplies the "pass it along" half through
embedding; each concrete handler supplies only its own condition.

## Task

Implement `Handle` on `*H20` in [chainresp.go](chainresp.go):

1. If `req == 20`, return `"twenty"`.
2. Otherwise return `h.Next(req)`.
3. `H10` already shows the shape — follow it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  h10.Handle(10)
Output: "ten"
```

**Example 2:**

```
Input:  h10.Handle(20)   // h10 forwards to h20
Output: "twenty"
```

**Example 3:**

```
Input:  h10.Handle(30)
Output: "unhandled"
```

_Explanation:_ nothing matched, and `h20.next` is nil, so `BaseHandler.Next` returns the fallback.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Struct embedding** | `H20{BaseHandler}` promotes `SetNext` and `Next` — `H20` gets them for free. |
| 2 | **Interface satisfaction via promotion** | `*H20` satisfies `Handler` because `Handle` is its own and `SetNext` is promoted. |
| 3 | **nil-terminated chain** | The last link's `next` is nil, which is where the fallback string comes from. |

## Hint

Three lines. The fallback is not your job — `h.Next(req)` already handles both
"there is a next link" and "there is not".

## Validate

```bash
make verify
```
