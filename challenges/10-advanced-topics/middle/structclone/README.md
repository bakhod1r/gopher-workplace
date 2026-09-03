# Copy A Struct Without Knowing Its Type

**Level:** middle
**Topic:** 10-advanced-topics / 03-reflection

## Context

A snapshot helper stores a copy of whatever struct it is handed. Type-asserting through a list of known types covered four of them and missed the fifth.

## Task

Implement [structclone.go](structclone.go):

1. Return a copy of the struct `v`, with the same dynamic type.
2. The copy is shallow — reference fields stay shared.
3. Return nil for a nil interface or anything that is not a struct.

Replace the stub body in [structclone.go](structclone.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Clone(pt{1,2})
Output: pt{1,2}
```

_Explanation:_ Same type, distinct value.

**Example 2:**

```
Input:  out := Clone(in).(pt); out.X = 99
Output: in.X unchanged
```

_Explanation:_ Value fields are copied.

**Example 3:**

```
Input:  Clone(&pt{})
Output: <nil>
```

_Explanation:_ A pointer is not a struct.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.New** | Allocates a new value of a type and returns a pointer Value to it. |
| 2 | **Value.Set** | Struct assignment through reflection is the same shallow copy as `=`. |
| 3 | **Value.Interface** | Boxes the reflected value back into an `any` with its real type. |

## Hint

`reflect.New(t)` gives a pointer. You want what it points at.

## Validate

```bash
make verify
```
