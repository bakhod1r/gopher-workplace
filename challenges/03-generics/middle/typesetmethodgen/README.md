# Type Set And Method

**Level:** middle  
**Topic:** 03-generics

## Context

Status codes are named integer types with a `String` method. A report needs both their labels and their arithmetic sum.

## Task

Implement the stub(s) in [typesetmethodgen.go](typesetmethodgen.go):

1. Implement `Labels`, returning the labels and the sum of the numeric values.
2. The `Code` constraint requires both an underlying `int` and a `String` method — use both.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Labels([]Status{200, 404})
Output: ["ok", "missing"], 604
```

**Example 2:**

```
Input:  Labels([]Status{})
Output: [], 0
```

**Example 3:**

```
Input:  Labels([]Status{200})
Output: ["ok"], 200
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mixed constraints** | `interface{ ~int; String() string }` demands an underlying type *and* a method. |
| 2 | **Both capabilities usable** | `int(v)` works because of the type set; `v.String()` because of the method. |
| 3 | **Narrow by design** | Only named int types with a `String` method can instantiate this. |

## Hint

The type set gives you `int(v)`; the method set gives you `v.String()`.

## Validate

```bash
make verify
```
