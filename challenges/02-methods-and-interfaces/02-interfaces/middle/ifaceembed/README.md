# Embedded Interface

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A monitoring wrapper adds counting to any metric source without reimplementing the source.

## Task

Implement the stub(s) in [ifaceembed.go](ifaceembed.go):

1. Implement `Value` on `Gauge`.
2. Implement `Value` on `*CountingSource` so it delegates to the embedded `Source` and increments `Calls`.
3. Implement `ReadTwice`, which reads the source twice and returns both values.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Gauge{N: 5}.Value()
Output: 5
```

**Example 2:**

```
Input:  c := &CountingSource{Source: Gauge{N: 5}}; c.Value(); c.Calls
Output: 1
```

**Example 3:**

```
Input:  ReadTwice(c)
Output: 5, 5 and c.Calls == 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Embedding an interface in a struct** | The struct inherits the interface's method set and can override it. |
| 2 | **Delegation** | The wrapper forwards to the embedded value, adding behaviour around it. |
| 3 | **Method promotion** | Reused from methods: embedded methods are promoted to the outer type. |

## Hint

Inside the override, call `c.Source.Value()` — not `c.Value()`, which would recurse.

## Validate

```bash
make verify
```
