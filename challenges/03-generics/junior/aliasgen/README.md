# Generic Type Alias

**Level:** junior  
**Topic:** 03-generics

## Context

The codebase writes `map[string]struct{}` everywhere. A generic alias gives that shape one readable name without inventing a new type.

## Task

Implement the stub(s) in [aliasgen.go](aliasgen.go):

1. Implement `NewIndex`, `Mark`, and `Marked`.
2. `Index[K]` is a generic **alias** for `map[K]struct{}` — the two are the same type, not merely convertible.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewIndex[string]()
Output: an empty Index
```

**Example 2:**

```
Input:  Mark(ix, "a"); Marked(ix, "a")
Output: true
```

**Example 3:**

```
Input:  Marked(ix, "b")
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generic type aliases** | `type Index[K comparable] = map[K]struct{}` names a shape without defining a new type. |
| 2 | **Alias versus definition** | An alias is interchangeable with its target; a defined type is not, and can carry methods. |
| 3 | **No methods on aliases** | Because it is not a new type, you cannot attach methods — hence the plain functions here. |

## Hint

`Index[string]` and `map[string]struct{}` are literally the same type — the tests rely on that.

## Validate

```bash
make verify
```
