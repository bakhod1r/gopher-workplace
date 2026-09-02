# Constraint With A Method

**Level:** middle  
**Topic:** 03-generics

## Context

Sensor readings come in several concrete types, each knowing how to report itself as a number.

## Task

Implement the stub(s) in [measurablegen.go](measurablegen.go):

1. Implement `Heaviest`, returning the element with the largest `Value()` and `true`.
2. On a tie keep the earlier element; return zero and `false` for an empty slice.
3. Call `Value()` at most once per element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Heaviest([]Reading{{1}, {5}})
Output: {5}, true
```

**Example 2:**

```
Input:  Heaviest([]Reading{{2}, {2}})
Output: the first, true
```

**Example 3:**

```
Input:  Heaviest([]Reading{})
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method constraints** | A constraint may require methods as well as underlying types. |
| 2 | **Typed slices stay typed** | The caller keeps `[]Reading`; no boxing into `[]Measurable` happens. |
| 3 | **Caching the projection** | `Value()` may be expensive — call it once per element. |

## Hint

The constraint is an ordinary interface with one method.

## Validate

```bash
make verify
```
