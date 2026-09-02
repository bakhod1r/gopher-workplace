# Flyer

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An airfield simulator asks each craft for its cruising altitude and picks the highest flyer.

## Task

Implement the stub(s) in [flyer.go](flyer.go):

1. Implement `Altitude` on `Bird` and `Jet`.
2. Implement `Highest`, which returns the greatest altitude in the slice, or 0 when empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Bird{Meters: 100}.Altitude()
Output: 100
```

**Example 2:**

```
Input:  Jet{Feet: 30000}.Altitude()
Output: 9144
```

**Example 3:**

```
Input:  Highest([]Flyer{Bird{Meters: 100}, Jet{Feet: 1000}})
Output: 304
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface as a unit adapter** | Both types report metres even though `Jet` stores feet. |
| 2 | **Integer division** | Reused from data types: `feet * 3048 / 10000` keeps it in ints. |
| 3 | **Max over a slice** | Reused: track the running maximum. |

## Hint

1 foot = 0.3048 m. Use `f.Feet * 3048 / 10000` to stay in integer arithmetic.

## Validate

```bash
make verify
```
