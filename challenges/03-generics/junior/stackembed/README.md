# Embedding A Generic Type

**Level:** junior  
**Topic:** 03-generics

## Context

An instrumented stack should count pushes without reimplementing the stack itself.

## Task

Implement the stub(s) in [stackembed.go](stackembed.go):

1. Implement `Push` on `TracedStack`, delegating to the embedded `Stack` and recording the value and count.
2. Implement `Pushes` and `Last`.
3. `Stack.Push` and `Stack.Len` are provided — do not change them.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(1); Push(2); Pushes()
Output: 2
```

**Example 2:**

```
Input:  Push(1); Push(2); Last()
Output: 2, true
```

**Example 3:**

```
Input:  Len() after two pushes
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Embedding an instantiated generic type** | `Stack[T]` is embedded inside `TracedStack[T]`, forwarding the parameter. |
| 2 | **Shadowing a promoted method** | `TracedStack.Push` hides `Stack.Push`, so the delegation must be explicit. |
| 3 | **Promotion still applies** | `Len` is promoted unchanged and needs no code. |

## Hint

Call `t.Stack.Push(v)` explicitly — your `Push` shadows the promoted one.

## Validate

```bash
make verify
```
