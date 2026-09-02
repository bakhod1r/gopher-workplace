# comparable Versus any

**Level:** middle  
**Topic:** 03-generics

## Context

A lookup keyed by `any` panicked in production when someone passed a slice as a key. The generic version cannot compile with such a key at all.

## Task

Implement the stub(s) in [comparablekeygen.go](comparablekeygen.go):

1. Implement `Index`, pairing keys with values positionally.
2. Stop at the shorter slice; later duplicate keys win.
3. `IndexAny` is provided — compare where each version fails.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Index([]string{"a"}, []int{1})
Output: {a:1}
```

**Example 2:**

```
Input:  Index with a slice key
Output: does not compile
```

**Example 3:**

```
Input:  IndexAny with a slice key
Output: panics at run time
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`comparable` is checked at compile time** | An uncomparable type argument is rejected at instantiation. |
| 2 | **`any` keys are checked at run time** | `map[any]V` panics when a stored key turns out to be uncomparable. |
| 3 | **Moving failures earlier** | This is the practical payoff of the constraint. |

## Hint

The body is a plain zip; the value is in what the constraint refuses.

## Validate

```bash
make verify
```
