# A Copy That Shares Nothing

**Level:** staff
**Topic:** 10-advanced-topics / 03-reflection

## Context

A cache hands out the struct it stores. Callers mutate a slice three levels down, the cached entry changes with it, and the bug is reported as "the cache returns wrong data at random".

## Task

Implement [deepcopy.go](deepcopy.go):

1. Return a copy of `v` that shares no mutable storage with it.
2. Recurse through structs, pointers, interfaces, slices, arrays and maps.
3. Preserve nil slices, maps and pointers as nil; skip unexported fields.
4. Scalars, strings and a nil interface are returned as they are.

Replace the stub body in [deepcopy.go](deepcopy.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  out := DeepCopy(in).(node); out.Tags[0] = "x"
Output: in.Tags is unchanged
```

**Example 2:**

```
Input:  out.Child == in.Child
Output: false
```

_Explanation:_ The pointed-at struct is copied, not the pointer.

**Example 3:**

```
Input:  DeepCopy(node{Name:"bare"})
Output: nil fields stay nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.MakeSlice / MakeMapWithSize** | Building new containers of a type known only at run time. |
| 2 | **Nil is not empty** | Copying a nil map into a made map changes observable behaviour. |
| 3 | **Recursion over kinds** | Each container kind needs its own construction step. |
| 4 | **Unexported fields are uncopyable** | `Set` refuses them, so they must be skipped. |

## Hint

The recursive helper is written for you. The exported function validates, recurses and boxes.

## Validate

```bash
make verify
```
