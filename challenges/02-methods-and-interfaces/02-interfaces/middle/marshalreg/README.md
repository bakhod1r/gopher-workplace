# Codec Registry

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A message bus picks an encoder by content type and falls back when the type is unknown.

## Task

Implement the stub(s) in [marshalreg.go](marshalreg.go):

1. Implement `Encode` on `CSVCodec` (join with `,`) and `PipeCodec` (join with `|`).
2. Implement `Register` and `Encode` on `*Registry`; an unknown type falls back to the registered default.
3. Implement `SetDefault`; with no default registered, `Encode` returns `ErrNoCodec`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  reg.Encode("csv", []string{"a", "b"})
Output: "a,b", nil
```

**Example 2:**

```
Input:  reg.Encode("unknown", ...) with a default of PipeCodec
Output: "a|b", nil
```

**Example 3:**

```
Input:  reg.Encode("unknown", ...) with no default
Output: "", ErrNoCodec
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Registry with a fallback** | Lookup, then default, then error — a common resolution order. |
| 2 | **Interface field** | The default is stored as an interface value, so it can be nil. |
| 3 | **strings.Join** | Reused from standard library: joining with a separator. |

## Hint

A nil interface field is the signal that no default was set — check it with `== nil`.

## Validate

```bash
make verify
```
