# Interface Composition

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A storage layer needs handles that can both read and write; some backends only do one.

## Task

Implement the stub(s) in [ifacecomp.go](ifacecomp.go):

1. Implement `Read`/`Write` on `*File` so it satisfies `ReadWriter`.
2. Implement `Describe`, which reports `"rw"`, `"r"`, `"w"`, or `"none"` depending on which interfaces the value satisfies.
3. Use interface-to-interface assertions, not concrete types.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Describe(&File{})
Output: "rw"
```

**Example 2:**

```
Input:  Describe(ReadOnly{Data: "x"})
Output: "r"
```

**Example 3:**

```
Input:  Describe(struct{}{})
Output: "none"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface embedding** | `ReadWriter` embeds `Reader` and `Writer` — the method set is the union. |
| 2 | **Interface-to-interface assertion** | `v.(Reader)` succeeds when the dynamic type has the method. |
| 3 | **Method sets** | Reused: pointer receivers mean only `*File` satisfies the interface. |

## Hint

Assert against `Reader` and `Writer` separately, then combine the two bools.

## Validate

```bash
make verify
```
