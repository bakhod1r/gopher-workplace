# Config Value

**Level:** junior
**Topic:** 04-error-handling

## Context

A service reads settings from a string map at startup. A missing key must stop the boot rather than run with an empty value.

## Task

Implement `Get` in [cfgget.go](cfgget.go):

1. Return the value stored under `key` and nil.
2. Return `""` and `ErrMissingKey` when the key is absent.
3. An empty stored value is present, not missing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Get(map[string]string{"port": "80"}, "port")
Output: "80", nil
```

**Example 2:**

```
Input:  Get(map[string]string{"a": ""}, "a")
Output: "", nil
```

**Example 3:**

```
Input:  Get(nil, "port")
Output: "", ErrMissingKey
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok idiom** | Presence and value are separate answers. |
| 2 | **Empty vs missing** | A stored empty string is a real configured value. |
| 3 | **Fail fast at startup** | Configuration errors surface before traffic. |

## Hint

One of the test cases stores an empty string on purpose — presence is decided by `ok`, not by the value.

## Validate

```bash
make verify
```
