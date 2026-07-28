# Map Get With Default

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Reading a missing map key returns the zero value — indistinguishable from a real
zero. The comma-ok idiom tells them apart.

## Task

Implement `GetOr(m, key, def)` returning the value if present, else `def`.

## Examples

```go
GetOr({a:1}, "a", 99)      // => 1
GetOr({zero:0}, "zero", 99)// => 0  (present, not default)
GetOr({}, "x", 99)         // => 99
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok** | `v, ok := m[key]` reports presence. |
| 2 | **Zero vs absent** | A stored 0 differs from a missing key. |
| 3 | **Nil map read** | Reading nil is safe, returns absent. |

## Hint

`if v, ok := m[key]; ok { return v }; return def`.

## Validate

```bash
make verify
```
