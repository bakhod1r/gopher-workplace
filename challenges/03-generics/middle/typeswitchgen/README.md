# Type Switching On A Parameter

**Level:** middle  
**Topic:** 03-generics

## Context

A logger prints a hint about the value's kind. Someone tried to switch on the type parameter directly and the compiler refused.

## Task

Implement the stub(s) in [typeswitchgen.go](typeswitchgen.go):

1. Implement `Describe`, returning `"integer"`, `"float"`, `"string"`, or `"other"`.
2. A type parameter cannot be type-switched directly — convert to `any` first.
3. Named types whose underlying type is `int` report `"other"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Describe(1)
Output: "integer"
```

**Example 2:**

```
Input:  Describe("a")
Output: "string"
```

**Example 3:**

```
Input:  Describe(true)
Output: "other"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A real limit of Go generics** | Some things simply cannot be expressed — knowing which saves hours. |
| 2 | **`any(v)` is the bridge** | Boxing the value gives the switch a dynamic type to inspect. |
| 3 | **Dynamic type is exact** | A `type Celsius float64` matches neither `float64` nor `float32` in the switch. |

## Hint

`switch any(v).(type)` — the conversion is what makes it legal.

## Validate

```bash
make verify
```
