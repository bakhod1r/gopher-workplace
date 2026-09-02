# Marshaler

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An export job serialises records into a wire format the receiver understands.

## Task

Implement the stub(s) in [marshaler.go](marshaler.go):

1. Implement `Marshal` on `Point` — return `"<X>,<Y>"`.
2. Implement `Marshal` on `Label` — return the string itself.
3. Implement `MarshalAll`, which marshals every value in order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Point{X: 1, Y: 2}.Marshal()
Output: "1,2"
```

**Example 2:**

```
Input:  Label("hi").Marshal()
Output: "hi"
```

**Example 3:**

```
Input:  MarshalAll([]Marshaler{Point{X: 0, Y: 0}, Label("a")})
Output: ["0,0", "a"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Serialisation behind an interface** | Each type owns its own wire form. |
| 2 | **strconv.Itoa** | Reused from standard library basics: int to string. |
| 3 | **Defined string type** | `type Label string` satisfies the interface with one method. |

## Hint

`strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)`.

## Validate

```bash
make verify
```
