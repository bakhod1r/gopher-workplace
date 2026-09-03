# Build A Value From A Name

**Level:** staff
**Topic:** 10-advanced-topics / 03-reflection

## Context

A message broker decodes payloads into the struct named in the envelope. The dispatch is a switch over forty type names that has to be edited every time a message type is added.

## Task

Implement [typeregistry.go](typeregistry.go):

1. Return a pointer to a freshly allocated zero value of the type registered under `name`.
2. Every call returns a distinct value.
3. Return `ErrUnknown` when nothing is registered under that name.
4. Correct under concurrent use — many goroutines constructing at once.

Replace the stub body in [typeregistry.go](typeregistry.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  New("job")
Output: *job pointing at the zero job, nil
```

**Example 2:**

```
Input:  a, _ := New("job"); b, _ := New("job")
Output: a != b
```

_Explanation:_ Each call allocates its own value.

**Example 3:**

```
Input:  New("nope")
Output: ErrUnknown
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.New** | Allocates a zero value of a run-time type and returns a pointer Value to it. |
| 2 | **Value.Interface** | Boxes the pointer back with its real dynamic type, so `v.(*job)` succeeds. |
| 3 | **Types are immutable and shareable** | A `reflect.Type` is safe to hold and read from any goroutine. |
| 4 | **sync.Map for a read-mostly registry** | Written at init, read on every message. |

## Hint

`Register` and `lookup` are given. Three lines: look up, guard, construct.

## Validate

```bash
make verify
```
