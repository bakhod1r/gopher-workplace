# Behavioural Interface

**Level:** middle
**Topic:** 04-error-handling

## Context

A client retries only failures the transport marked as temporary. The concrete type does not matter — only whether it advertises the behaviour.

## Task

Implement `IsTemporary` in [temperr.go](temperr.go):

1. Return `true` when the chain contains an error with a `Temporary() bool` method returning true.
2. Return `false` when the method exists but returns false.
3. Return `false` when no error in the chain has the method.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IsTemporary(&NetError{Temp: true})
Output: true
```

**Example 2:**

```
Input:  IsTemporary(&NetError{Temp: false})
Output: false
```

**Example 3:**

```
Input:  IsTemporary(errors.New("boom"))
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Behavioural interfaces** | Match on what an error can do, not what it is. |
| 2 | **errors.As with an interface target** | `As` accepts an interface pointer. |
| 3 | **Chain traversal** | The behaviour may be several layers down. |

## Hint

Declare a local interface with the single method, then let `errors.As` find anything satisfying it.

## Validate

```bash
make verify
```
