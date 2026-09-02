# Type Switch Dispatch

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A hot decoder used a long type switch. Replacing it with a map of handlers changed both the performance profile and the extensibility story.

## Task

Implement the stub(s) in [typeswitchtab.go](typeswitchtab.go):

1. Implement `DecodeSwitch`, which classifies a value with a type switch.
2. Implement `Register` and `DecodeTable` on `*Table`, dispatching by `reflect.Type` through a map.
3. Both must return the same labels: `"int"`, `"string"`, `"bool"`, or `"unknown"`.
4. Constraint: `DecodeSwitch` must allocate zero times per call, and the table lookup must not depend on registration order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DecodeSwitch(1)
Output: "int"
```

**Example 2:**

```
Input:  DecodeTable after registering int and string
Output: the same labels
```

**Example 3:**

```
Input:  an unregistered type
Output: "unknown"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type switch lowering** | A type switch compares type descriptors; long chains cost more than a hash lookup. |
| 2 | **reflect.Type as a map key** | Type descriptors are comparable and canonical, so they key a dispatch table. |
| 3 | **Extensibility trade-off** | A switch is closed at compile time; a table is open at run time. |

## Hint

`reflect.TypeOf(v)` is nil for a nil interface — handle that before the lookup.

## Validate

```bash
make verify
```
