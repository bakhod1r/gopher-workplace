# Methods Take No Type Parameters

**Level:** middle  
**Topic:** 03-generics

## Context

The team wants `bag.Map(f)` to sit next to `bag.Add(v)`. Go allows one and not the other, and the reason shapes every generic API in the language.

## Task

Implement the stub(s) in [methodtypeparamgen.go](methodtypeparamgen.go):

1. Implement `MapBag` as a function, `Add` and `Items` as methods.
2. `Add` returns a new bag rather than mutating the receiver.
3. Note that `Map` cannot be a method: it needs a type parameter of its own.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Bag[int]{}.Add(1).Add(2).Items()
Output: [1 2]
```

**Example 2:**

```
Input:  MapBag(bag, itoa).Items()
Output: ["1" "2"]
```

**Example 3:**

```
Input:  MapBag(empty, itoa).Items()
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A real limit of Go generics** | Some things simply cannot be expressed — knowing which saves hours. |
| 2 | **Where new parameters may live** | Functions may introduce type parameters; methods may only use the type's. |
| 3 | **Value semantics** | Returning a new bag keeps `Add` chainable without a pointer receiver. |

## Hint

`Add` reuses `T`; `Map` would introduce `U`, so it must be a free function.

## Validate

```bash
make verify
```
