# errors.Is

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A storage client wraps low-level failures but callers still need to recognise the sentinel underneath.

## Task

Implement the stub(s) in [errorsis.go](errorsis.go):

1. Implement `Fetch`, which returns `ErrNotFound` wrapped with the key when the key is missing.
2. Implement `IsMissing`, which reports whether an error is (or wraps) `ErrNotFound`.
3. Implement `FetchAll`, which returns the values for all keys, or the first wrapped error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Fetch(map[string]string{}, "k")
Output: error "fetch k: not found"
```

**Example 2:**

```
Input:  IsMissing(that error)
Output: true
```

**Example 3:**

```
Input:  FetchAll(data, []string{"a"}) with data["a"]="1"
Output: ["1"], nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Error wrapping with %w** | `fmt.Errorf("...: %w", err)` keeps the chain inspectable. |
| 2 | **errors.Is** | Compares against every error in the chain, not just the outermost. |
| 3 | **Sentinel errors** | Reused from error handling: package-level comparable values. |

## Hint

`%w` wraps; `%v` flattens to text and breaks `errors.Is`.

## Validate

```bash
make verify
```
