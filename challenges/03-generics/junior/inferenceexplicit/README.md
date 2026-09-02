# When Inference Fails

**Level:** junior  
**Topic:** 03-generics

## Context

A helper package keeps tripping people up: some calls need `[int]` spelled out and some do not, and nobody can say why.

## Task

Implement the stub(s) in [inferenceexplicit.go](inferenceexplicit.go):

1. Implement `Empty`, `ZeroOf`, and `Wrap`.
2. Note which calls need an explicit type argument and which infer it.
3. Do not change the signatures.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Empty[int]()
Output: []int{}
```

**Example 2:**

```
Input:  ZeroOf[string]()
Output: ""
```

**Example 3:**

```
Input:  Wrap(5)
Output: []int{5}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Inference comes from arguments** | No arguments means nothing to infer from — the type must be written. |
| 2 | **Return types never drive inference** | Go does not infer a type parameter from the assignment target. |
| 3 | **Explicit instantiation** | `Empty[int]()` names the type argument at the call site. |

## Hint

`Wrap(5)` infers; `Empty()` cannot — there is no argument to look at.

## Validate

```bash
make verify
```
