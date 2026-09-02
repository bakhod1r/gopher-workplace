# Any Value

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A metrics bag stores untyped values and must answer typed queries about them safely.

## Task

Implement the stub(s) in [anyvalue.go](anyvalue.go):

1. Implement `Set` and `Len` on `*Bag`.
2. Implement `GetString`, which returns the string at the key and whether it was both present and a string.
3. Implement `Kinds`, which returns the sorted list of distinct dynamic type names in the bag (`"int"`, `"string"`, `"bool"`, `"other"`).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  b.Set("a", "x"); b.GetString("a")
Output: "x", true
```

**Example 2:**

```
Input:  b.Set("n", 1); b.GetString("n")
Output: "", false
```

**Example 3:**

```
Input:  Kinds after Set("a","x"), Set("n",1)
Output: ["int", "string"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`any` as a value box** | `map[string]any` stores heterogeneous data. |
| 2 | **Assertion vs presence** | Two things can fail: the key is absent, or the type is wrong. |
| 3 | **sort.Strings** | Reused from standard library: deterministic output order. |

## Hint

`v, ok := b.data[key]` first, then `s, ok := v.(string)` — both must succeed.

## Validate

```bash
make verify
```
