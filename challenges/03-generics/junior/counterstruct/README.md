# Generic Counter

**Level:** junior  
**Topic:** 03-generics

## Context

A log analyser tallies status codes; a word tool tallies words. Both want the same counting structure.

## Task

Implement the stub(s) in [counterstruct.go](counterstruct.go):

1. Implement `NewCounter`, `Inc`, `Count`, and `Total`.
2. `Count` returns `0` for a value that was never incremented.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Inc("a"); Inc("a"); Count("a")
Output: 2
```

**Example 2:**

```
Input:  Count("never")
Output: 0
```

**Example 3:**

```
Input:  Inc("a"); Inc("b"); Total()
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map keys need `comparable`** | A generic type storing a map must constrain its key parameter. |
| 2 | **Missing keys read as zero** | Reused from language basics: `c.counts[v]` on an absent key is `0`, which is exactly the right tally. |
| 3 | **Receivers repeat the parameter** | A method on `Stack[T]` writes the receiver as `(s *Stack[T])` — the parameter comes along. |

## Hint

`c.counts[v]++` works even for a key that is not there yet.

## Validate

```bash
make verify
```
