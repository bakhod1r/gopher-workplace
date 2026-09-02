# Safe Publication

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A config pointer was published with a plain store. On a weakly ordered machine readers observed the pointer before the object it points to was fully initialised.

## Task

Implement the stub(s) in [memmodelpub.go](memmodelpub.go):

1. Implement `Publish` and `Load` on `*Publisher` using `atomic.Pointer` so the object's initialising writes happen-before every successful load.
2. Implement `Ready`, reporting whether anything has been published.
3. Constraint: `-race` clean with concurrent publishers and readers, and a loaded object must never be observed partially initialised.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Publish then Load
Output: the fully initialised object
```

**Example 2:**

```
Input:  Load before any Publish
Output: nil, false
```

**Example 3:**

```
Input:  concurrent publishers and readers
Output: every observed object is complete
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Go memory model** | An atomic store followed by an atomic load of the same location establishes happens-before. |
| 2 | **Safe publication** | Initialise fully, then publish the pointer — never the other way round. |
| 3 | **atomic.Pointer** | Typed atomic pointer without `unsafe` or interface boxing. |

## Hint

Build the object completely into a local, then store the pointer once. Never mutate a published object.

## Validate

```bash
make verify
```
