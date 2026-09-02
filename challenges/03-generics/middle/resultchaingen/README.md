# Chaining Results

**Level:** middle  
**Topic:** 03-generics

## Context

A validation pipeline runs several steps, each of which may fail. Once one fails the rest must be skipped, and the reason must survive to the end.

## Task

Implement the stub(s) in [resultchaingen.go](resultchaingen.go):

1. Implement `Ok`, `Fail`, `Then`, and `Unwrap`.
2. `Then` must be a function, not a method — it introduces the type parameter `U`.
3. A failure propagates unchanged, and `f` is not called.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Then(Ok(2), double)
Output: Ok(4)
```

**Example 2:**

```
Input:  Then(Fail[int]("bad"), double)
Output: failure, reason kept
```

**Example 3:**

```
Input:  f not called on failure
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Why `Then` is not a method** | It changes the payload type, and methods cannot add type parameters. |
| 2 | **Short-circuiting** | Skipping `f` on failure is what makes a chain safe to write linearly. |
| 3 | **Carrying the reason** | The failure must be rebuilt as `Result[U]`, keeping the original message. |

## Hint

`Then` cannot be a method — `U` is a new type parameter.

## Validate

```bash
make verify
```
