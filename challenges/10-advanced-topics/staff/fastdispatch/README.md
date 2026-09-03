# Reflect Only When You Have To

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A structured logger renders every field through reflection. Ninety-nine percent of the fields are ints and strings, and each one pays for a run-time type walk.

## Task

Implement [fastdispatch.go](fastdispatch.go):

1. Append `v`'s text form to `dst`.
2. Handle nil, `string`, `int`, `int64`, `bool` and `[]byte` with a type switch — no allocations.
3. Fall back to reflection for other integer, string and bool kinds, including named types.
4. Anything else renders as `?`.

Replace the stub body in [fastdispatch.go](fastdispatch.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Render(nil, 42)
Output: "42"
```

_Explanation:_ The fast path.

**Example 2:**

```
Input:  Render(nil, myInt(5))
Output: "5"
```

_Explanation:_ A named type reaches the fallback.

**Example 3:**

```
Input:  Render(nil, 1.5)
Output: "?"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type switch vs reflection** | The switch compares type words; reflection walks the type. |
| 2 | **Named types miss the switch** | `myInt` is not `int` to a type assertion, but its kind is int. |
| 3 | **strconv.Append*** | Renders into the destination with no intermediate string. |
| 4 | **Fast path, slow path** | Optimise for the distribution you actually have. |

## Hint

Two switches: one on the concrete type, one on the kind.

## Validate

```bash
make verify
```
