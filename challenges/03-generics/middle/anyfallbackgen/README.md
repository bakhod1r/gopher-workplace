# When Generics Are Not Enough

**Level:** middle  
**Topic:** 03-generics

## Context

A config loader receives decoded JSON of unknown shape. Generics cannot express "a slice nested to any depth", so the honest answer is an interface.

## Task

Implement the stub(s) in [anyfallbackgen.go](anyfallbackgen.go):

1. Implement `DeepCount`, counting leaf values inside nested `[]any` values.
2. A `nil` counts as zero; any non-slice value counts as one leaf.
3. Note why a type parameter cannot express this shape.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DeepCount([]any{1, 2})
Output: 2
```

**Example 2:**

```
Input:  DeepCount([]any{1, []any{2, 3}})
Output: 3
```

**Example 3:**

```
Input:  DeepCount(nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Falling back to `any`** | When the type set cannot express it, an interface plus a type switch still can. |
| 2 | **Unbounded nesting** | `[]T`, `[][]T`, `[][][]T` are different types; no constraint spans them all. |
| 3 | **Recursive type switches** | Each level is inspected dynamically, because the depth is not known statically. |

## Hint

Generics fix the depth at compile time; this shape needs a dynamic walk.

## Validate

```bash
make verify
```
