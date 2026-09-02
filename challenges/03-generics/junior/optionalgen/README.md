# Optional

**Level:** junior  
**Topic:** 03-generics

## Context

A parser returns "no value here" for optional fields. Callers should not have to juggle pointers to express that.

## Task

Implement the stub(s) in [optionalgen.go](optionalgen.go):

1. Implement `Some`, returning an `Optional` holding `v`.
2. Implement `None`, returning an empty `Optional`.
3. Implement `Or`, returning the held value or `def`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Some(5).Or(9)
Output: 5
```

**Example 2:**

```
Input:  None[int]().Or(9)
Output: 9
```

**Example 3:**

```
Input:  Some(0).Or(9)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Constructor functions** | `Some` and `None` are plain generic functions returning an instantiated type. |
| 2 | **Explicit instantiation** | `None[int]()` must name `T`: there is no argument to infer from. |
| 3 | **Value receivers** | `Or` only reads, so a value receiver is right. |

## Hint

`None[int]()` needs its type argument written out.

## Validate

```bash
make verify
```
