# Bounded Cache

**Level:** junior  
**Topic:** 03-generics

## Context

A lookup service caches recent results but must not grow without bound.

## Task

Implement the stub(s) in [cachegen.go](cachegen.go):

1. Implement `NewCache`, `Put`, and `Get`.
2. `Put` evicts the oldest entry once the cache is over capacity.
3. Re-putting an existing key updates its value without changing its insertion order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewCache[string,int](2); Put(a,1); Put(b,2); Put(c,3); Get(a)
Output: 0, false
```

**Example 2:**

```
Input:  Put(a,1); Put(a,2); Get(a)
Output: 2, true
```

**Example 3:**

```
Input:  NewCache[string,int](0); Put(a,1); Get(a)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two parallel structures** | The map answers lookups; the slice remembers insertion order. |
| 2 | **Eviction policy** | Dropping `order[0]` evicts the oldest inserted key — a first-in-first-out cache. |
| 3 | **Map keys need `comparable`** | A generic type storing a map must constrain its key parameter. |

## Hint

Only append to `order` when the key is new — otherwise updates would evict too early.

## Validate

```bash
make verify
```
