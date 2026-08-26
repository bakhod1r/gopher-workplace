# Promote

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A widget system embeds a `Base` component. The base provides `Hello()`, and
the extended widget uses it in its own `Describe()` method.

## Task

Implement `Describe` on `Extended` in [promote.go](promote.go):

1. Call the promoted `Hello()` method (from embedded `Base`).
2. Return `"<Hello()> [<Extra>]"`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Extended{Base{"Go"}, "fast"}.Describe()
Output: "Hello from Go [fast]"
```

**Example 2:**

```
Input:  Extended{Base{"X"}, ""}.Describe()
Output: "Hello from X []"
```

**Example 3:**

```
Input:  Extended{Base{""}, "y"}.Describe()
Output: "Hello from  [y]"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Struct embedding** | `Extended` embeds `Base` — `Base`'s methods are promoted. |
| 2 | **Method promotion** | `e.Hello()` calls `Base.Hello()` on the embedded field. |

## Hint

`e.Hello() + " [" + e.Extra + "]"` — `Hello` is promoted from `Base`.

## Validate

```bash
make verify
```
