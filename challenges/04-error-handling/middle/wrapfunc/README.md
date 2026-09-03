# Decorate A Step

**Level:** middle
**Topic:** 04-error-handling

## Context

A pipeline wraps each stage so its failures are labelled without every stage having to remember its own name.

## Task

Implement `Named` in [wrapfunc.go](wrapfunc.go):

1. Return a function that calls `f` and returns its result.
2. Annotate a non-nil result as `"<name>: <err>"`, preserving the chain.
3. Pass a nil result through unchanged.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Named("load", okFn)()
Output: nil
```

**Example 2:**

```
Input:  Named("load", failFn)()
Output: "load: boom"
```

**Example 3:**

```
Input:  errors.Is(Named("load", failFn)(), ErrBoom)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closures** | The returned function captures `name` and `f`. |
| 2 | **Higher-order functions** | A decorator takes and returns a `func() error`. |
| 3 | **Deferred wrapping** | Annotation happens when the step runs, not when it is built. |

## Hint

Building the decorator must not call `f` — the tests check that nothing runs until the returned function is invoked.

## Validate

```bash
make verify
```
