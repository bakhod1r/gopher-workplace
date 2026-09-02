# Unit Conversion

**Level:** middle  
**Topic:** 03-generics

## Context

Temperatures are stored as `Celsius`; a third-party formula works in plain `float64`. The boundary must convert both ways without losing the unit.

## Task

Implement the stub(s) in [tempconv.go](tempconv.go):

1. Implement `ToFloat`, `FromFloat`, and `Rescale`.
2. `Rescale` must return the same named type it received.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ToFloat(Celsius(20))
Output: 20.0
```

**Example 2:**

```
Input:  FromFloat[Celsius](20)
Output: Celsius(20)
```

**Example 3:**

```
Input:  Rescale(Celsius(20), double)
Output: Celsius(40)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Converting a type parameter** | `float64(v)` and `T(f)` are both legal when the set is `~float64`. |
| 2 | **Explicit instantiation** | `FromFloat[Celsius](20)` must name `T` — the argument is a plain `float64`. |
| 3 | **Round-tripping** | Convert out, compute, convert back: the unit survives the third-party call. |

## Hint

`FromFloat` needs its type argument spelled out at the call site.

## Validate

```bash
make verify
```
