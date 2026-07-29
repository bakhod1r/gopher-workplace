# Value Receiver Loses Mutation

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A value receiver `(c Counter)` operates on a COPY; `c.N++` changes the copy and
vanishes when the method returns. Mutating methods need a pointer receiver.

## Task

Fix the receiver in [valuerecv.go](valuerecv.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := &Counter{}; c.Inc(); c.Inc()
Output: c.N == 2
```

**Example 2:**

```
Input:  c := &Counter{N: 5}; c.Inc()
Output: c.N == 6
```

**Example 3:**

```
Input:  c := &Counter{}; c.Inc()
Output: c.N == 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Value vs pointer receiver** | Value receiver mutates a copy. |
| 2 | **Persisting mutation** | Use `(c *Counter)`. |
| 3 | **Method sets** | Pointer receivers can mutate the callee. |

## Hint

Change the receiver to a pointer: `func (c *Counter) Inc()`.

## Validate

```bash
make verify
```
